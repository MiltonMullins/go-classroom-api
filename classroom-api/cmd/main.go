package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/miltonmullins/classroom-api/classroom-api/cmd/middleware"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/handlers"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/repositories"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/services"
)

func main() {
	//connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table participants if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS participants (
		id SERIAL PRIMARY KEY,
		classroom_id int references classrooms(id),
		user_id int references users(id),
		role VARCHAR(100) NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	router := initializeRoutes(db) // configure routes

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening on port:8080...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(db *sql.DB) *http.ServeMux {
	participantsHandler := handlers.NewParticipantsHandler(
		services.NewParticipantsService(
			repositories.NewParticipantsRepository(db)))

	mux := http.NewServeMux()
	mux.Handle("GET /participants", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(participantsHandler.GetParticipantsByClassroomID))))
	mux.Handle("POST /participants", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(participantsHandler.CreateParticipant))))
	mux.Handle("DELETE /participants", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(participantsHandler.DeleteParticipant))))
	return mux
}
