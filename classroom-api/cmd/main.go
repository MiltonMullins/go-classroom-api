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

	//create the table classrooms if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS classrooms (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		teacher_id int NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	//create the table participants if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS participants (
		id SERIAL PRIMARY KEY,
		classroom_id int references classrooms(id),
		user_id int NOT NULL,
		role VARCHAR(100) NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
	}

	router := initializeRoutes(db) // configure routes

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Println("Listening on port:8081...")
	server.ListenAndServe() // Run the http server
}

func initializeRoutes(db *sql.DB) *http.ServeMux {
	participantsHandler := handlers.NewParticipantsHandler(
		services.NewParticipantsService(
			repositories.NewParticipantsRepository(db)))

	classroomsHandler := handlers.NewClassroomHandler(
		services.NewClassroomService(
			repositories.NewClassroomRepository(db)))

	mux := http.NewServeMux()

	//Participants routes
	mux.Handle("GET /participants/{classroomID}", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(participantsHandler.GetParticipantsByClassroomID))))
	mux.Handle("POST /participant", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(participantsHandler.CreateParticipant))))
	mux.Handle("DELETE /participant/{classroomID}/{id}", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(participantsHandler.DeleteParticipant))))

	//Classrooms routes
	mux.Handle("GET /classrooms", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(classroomsHandler.GetClassrooms))))
	mux.Handle("GET /classroom/{id}", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(classroomsHandler.GetClassroomByID))))
	mux.Handle("POST /classrooms", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(classroomsHandler.CreateClassroom))))
	mux.Handle("PUT /classroom/{id}", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(classroomsHandler.UpdateClassroom))))
	mux.Handle("DELETE /classroom/{id}", middleware.JwtAuthMiddleware(middleware.Log(http.HandlerFunc(classroomsHandler.DeleteClassroom))))

	return mux
}
