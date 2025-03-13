package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/miltonmullins/classroom-api/notifications-api/internal/models"
)

func main() {

	router := inicializateRoutes()
	server := &http.Server{
		Addr:    ":8084",
		Handler: router,
	}

	log.Println("Listening on port:8084...")
	server.ListenAndServe() // Run the http server
}

func inicializateRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /push", sendHandler)
	mux.HandleFunc("GET /get", getHandler)

	return mux
}

func redisConnection() *redis.Client {
	//Create client
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0, // use default DB
	})

	return client
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	var notification models.Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := redisConnection()
	key := "Notification"+ strconv.Itoa(notification.UserID)
	value := "Eroll to " + notification.AssigmentTitle + " successful"
	errSet := client.Set(key, value, 0).Err()
	if errSet != nil {
		http.Error(w, errSet.Error(), http.StatusInternalServerError)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	client := redisConnection()

	val, err := client.Get("Notification123").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val)) //TODO: check this

}
