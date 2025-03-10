package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/miltonmullins/classroom-api/users-api/internal/services"
	"github.com/miltonmullins/classroom-api/users-api/internal/models"
)

type LoginHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
}

type loginHandler struct {
	LoginService services.LoginService
}

func NewLoginHandler(loginService services.LoginService) LoginHandler {
	return &loginHandler{
		LoginService: loginService,
	}
}

func (l *loginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jwtToken := l.LoginService.Login(&user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Token: %s", jwtToken)
}

