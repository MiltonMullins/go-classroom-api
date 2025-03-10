package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/miltonmullins/classroom-api/users-api/cmd/jwt"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing JWT Auth")
		w.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]//TODO Handle panic if tokenString isn't in correct format (Bearer <token>)

		err := jwt.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Log(next http.Handler) http.Handler {
	log.Printf("Log: %v", time.Now())
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareLog")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("Executing middlewareLog again")
	})
}
