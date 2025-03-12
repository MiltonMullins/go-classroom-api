package main

import (
	"context"
	"net/http"
	"log"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/handlers"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/services"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/repositories"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	defer func() {
	 if err = client.Disconnect(ctx); err != nil {
	  panic(err)
	 }
	}()

	router := initializeRoutes(client) // configure routes

	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	log.Println("Listening on port:8082...")
	server.ListenAndServe() // Run the http server
}


func initializeRoutes(db *mongo.Client) *http.ServeMux {
	assigmentHandler := handlers.NewAssigmentHandler(
		services.NewAssigmentService(
			repositories.NewAssigmentRepository(db)))

	mux := http.NewServeMux()

	//CRUD
	mux.HandleFunc("GET /assigment/{param}", assigmentHandler.GetAssigment)
	mux.HandleFunc("POST /assigment", assigmentHandler.CreateAssigment)
	mux.HandleFunc("UPDATE /assigment/{id}", assigmentHandler.UpdateAssigment)
	mux.HandleFunc("DELETE /assigment/{id}", assigmentHandler.DeleteAssigment)
	return mux
}