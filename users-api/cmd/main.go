package cmd

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"net/http"

	"github.com/miltonmullins/classroom-api/users-api/internal/handlers"
	"github.com/miltonmullins/classroom-api/users-api/internal/repositories"
	"github.com/miltonmullins/classroom-api/users-api/internal/services"
)

func main() {
	
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS people (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		age INT NOT NULL
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
	handler := handlers.NewUserHandler(
		services.NewUserService(
			repositories.NewUserRepository(db)))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/{id}", handler.GetUserById)
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("POST /user", handler.CreateUser)
	mux.HandleFunc("PUT /user/{id}", handler.UpdateUser)
	mux.HandleFunc("DELETE /user/{id}", handler.DeleteUser)
	return mux
}
