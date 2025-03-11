package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miltonmullins/classroom-api/classroom-api/internal/models"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/services"
)

type ClassroomHandler interface {
	GetClassrooms(w http.ResponseWriter, r *http.Request)
	GetClassroomByID(w http.ResponseWriter, r *http.Request)
	CreateClassroom(w http.ResponseWriter, r *http.Request)
	UpdateClassroom(w http.ResponseWriter, r *http.Request)
	DeleteClassroom(w http.ResponseWriter, r *http.Request)
}

type classroomHandler struct {
	service services.ClassroomService
}

func NewClassroomHandler(service services.ClassroomService) ClassroomHandler {
	return &classroomHandler{
		service: service,
	}
}

func (h *classroomHandler) GetClassrooms(w http.ResponseWriter, r *http.Request) {
	classrooms, err := h.service.GetClassrooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonClassrooms, err := json.Marshal(classrooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonClassrooms)
}

func (h *classroomHandler) GetClassroomByID(w http.ResponseWriter, r *http.Request) {
	classroomID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	classroom, err := h.service.GetClassroomByID(classroomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonClassroom, err := json.Marshal(classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonClassroom)
}

func (h *classroomHandler) CreateClassroom(w http.ResponseWriter, r *http.Request) {
	classroom := &models.Classroom{}
	err := json.NewDecoder(r.Body).Decode(classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.CreateClassroom(classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *classroomHandler) UpdateClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	classroom := &models.Classroom{}
	err = json.NewDecoder(r.Body).Decode(classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.UpdateClassroom(classroomID, classroom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *classroomHandler) DeleteClassroom(w http.ResponseWriter, r *http.Request) {
	classroomID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteClassroom(classroomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}