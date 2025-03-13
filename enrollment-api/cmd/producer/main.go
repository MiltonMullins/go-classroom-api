package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/IBM/sarama"

	"github.com/miltonmullins/classroom-api/enroll-api/internal/models"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /enroll", PlaceEnrollMessage)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}
	log.Println("Listening on port:8083...")
	server.ListenAndServe()
}

func PlaceEnrollMessage(w http.ResponseWriter, r *http.Request) {
	// Parse Request body
	var enrollMessage models.EnrollMessage

	if err := json.NewDecoder(r.Body).Decode(&enrollMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Convert body into bytes
	msgInBytes, err := json.Marshal(enrollMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Send the bytes to kafka
	err = PushMessageToQueue("enroll", msgInBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Respond back to the user
	response := map[string]interface{}{
		"success": true,
		"msg":     "Message for" + enrollMessage.StudentID + " placed successfully",
	}

	w.Header().Set("Contetn-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		http.Error(w, "Error placin order", http.StatusInternalServerError)
		return
	}
}


func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	//config.Producer.Retry = {5, 20 *time.Secons}

	return sarama.NewSyncProducer(brokers, config)
}

func PushMessageToQueue(topic string, message []byte) error {
	brokers := []string{"kafka:9092"}

	//create connection
	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()

	//Create kafka msg
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	//Send message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}