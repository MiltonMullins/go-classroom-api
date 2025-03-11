package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/miltonmullins/classroom-api/users-api/utils"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing JWT Auth")
		//w.Header().Set("Content-Type", "application/json") necesario?
		
		authHeader := r.Header.Get("Authorization")
		tokenSplit := strings.Split(authHeader, " ")
		if len(tokenSplit) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid authorization header")
			return
		}

		authToken := tokenSplit[1]

		err := jwt.VerifyToken(authToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		email, err := jwt.ExtractIDFromToken(authToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		w.Header().Set("x-user-email", email)
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
