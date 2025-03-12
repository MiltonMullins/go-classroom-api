package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/handlers"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/repositories"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/services"
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

	studentTaskHandler := handlers.NewStudentTasklHandler(
		services.NerStudentTaskRepository(
			repositories.NewStudentTaskRepository(db)))

	mux := http.NewServeMux()
			//TODO MIDDLEWARE
	//CRUD Assigment
	mux.HandleFunc("GET /assigment/{title}", assigmentHandler.GetAssigment)
	mux.HandleFunc("POST /assigment", assigmentHandler.CreateAssigment)
	mux.HandleFunc("PUT /assigment/{title}", assigmentHandler.UpdateAssigment)
	mux.HandleFunc("DELETE /assigment/{title}", assigmentHandler.DeleteAssigment)

	//CRUD Student Task
	mux.HandleFunc("GET /student-task/{assigment_id}", studentTaskHandler.GetStudentTasks)
	mux.HandleFunc("POST /student-task", studentTaskHandler.CreateStudentTask)
	mux.HandleFunc("PUT /student-task/{student_id}/{assigment_id}", studentTaskHandler.UpdateStudentTask)
	mux.HandleFunc("DELETE /student-task/{student_id}/{assigment_id}", studentTaskHandler.DeleteStudentTask)
	return mux
}
