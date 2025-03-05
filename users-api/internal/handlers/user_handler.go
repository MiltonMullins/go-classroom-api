package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/miltonmullins/classroom-api/users-api/internal/models"
	"github.com/miltonmullins/classroom-api/users-api/internal/services"
)

type UserHandler interface {
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := u.userService.GetUserById(id)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}

func (u *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	jsonUsers, err := json.Marshal(u.userService.GetUsers())
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

func (u *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = u.userService.CreateUser(&user)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User Created")
}

func (u *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = u.userService.UpdateUser(id, &user)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Updated")
}

func (u *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.userService.DeleteUser(id)
	if err != nil {
		//TODO Handle Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User Deleted")
}
