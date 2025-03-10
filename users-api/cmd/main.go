package main

import (
	"database/sql"
	"log"
	"net/http"
	//"os"

	_ "github.com/lib/pq"

	"github.com/miltonmullins/classroom-api/users-api/internal/handlers"
	"github.com/miltonmullins/classroom-api/users-api/internal/repositories"
	"github.com/miltonmullins/classroom-api/users-api/internal/services"
	"github.com/miltonmullins/classroom-api/users-api/cmd/middleware"
)

func main() {

	connStr := "postgres://postgres:postgres@classroom_db:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create the table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		role VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL
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
	UserHandler := handlers.NewUserHandler(
		services.NewUserService(
			repositories.NewUserRepository(db)))

	loginHandler := handlers.NewLoginHandler(
		services.NewLoginService(
			repositories.NewUserRepository(db)))

	mux := http.NewServeMux()

	//Login
	mux.HandleFunc("POST /login", loginHandler.Login)
	//TODO: mux.HandleFunc("/change-password", loginHandler.ChangePassword)

	//CRUD
	mux.Handle("GET /user/{id}", middleware.JWTAuth(middleware.Log(http.HandlerFunc(UserHandler.GetUserById))))
	mux.HandleFunc("GET /users", UserHandler.GetUsers)
	mux.Handle("POST /user", middleware.JWTAuth(middleware.Log(http.HandlerFunc(UserHandler.CreateUser))))
	mux.Handle("PUT /user/{id}", middleware.JWTAuth(middleware.Log(http.HandlerFunc(UserHandler.UpdateUser))))
	mux.Handle("DELETE /user/{id}", middleware.JWTAuth(middleware.Log(http.HandlerFunc(UserHandler.DeleteUser))))
	return mux
}


