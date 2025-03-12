package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/models"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/services"
)

type AssigmentHandler interface {
	GetAssigment(w http.ResponseWriter, r *http.Request)
	CreateAssigment(w http.ResponseWriter, r *http.Request)
	UpdateAssigment(w http.ResponseWriter, r *http.Request)
	DeleteAssigment(w http.ResponseWriter, r *http.Request)
}

type assigmentHandler struct {
	assigmentService services.AssigmentService
}

func NewAssigmentHandler(assigmentService services.AssigmentService) AssigmentHandler {
	return &assigmentHandler{
		assigmentService: assigmentService,
	}
}

func (a *assigmentHandler) GetAssigment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	param := r.PathValue("param")
	assigments, err := a.assigmentService.GetAssigment(ctx, param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonAssigments, err := json.Marshal(assigments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAssigments)
}

func (a *assigmentHandler) CreateAssigment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var assigment models.Assigment
	err := json.NewDecoder(r.Body).Decode(&assigment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.assigmentService.CreateAssigment(ctx, &assigment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *assigmentHandler) UpdateAssigment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	title := r.PathValue("title")

	var assigment models.Assigment
	err := json.NewDecoder(r.Body).Decode(&assigment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = a.assigmentService.UpdateAssigment(ctx, title, &assigment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *assigmentHandler) DeleteAssigment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	title := r.PathValue("title")

	err := a.assigmentService.DeleteAssigment(ctx, title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
