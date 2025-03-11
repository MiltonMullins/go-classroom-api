package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miltonmullins/classroom-api/classroom-api/internal/services"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/models"
)

type ParticipantsHandler interface {
	GetParticipantsByClassroomID(w http.ResponseWriter, r *http.Request)
	CreateParticipant(w http.ResponseWriter, r *http.Request)
	DeleteParticipant(w http.ResponseWriter, r *http.Request)
}

type participantsHandler struct {
	service services.ParticipantsService
}

func NewParticipantsHandler(service services.ParticipantsService) ParticipantsHandler {
	return &participantsHandler{
		service: service,
	}
}

func (h *participantsHandler) GetParticipantsByClassroomID(w http.ResponseWriter, r *http.Request) {
	classroomID, err := strconv.Atoi(r.PathValue("classroomID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	participants, err := h.service.GetParticipantsByClassroomID(classroomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonParticipants, err := json.Marshal(participants)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonParticipants)
}

func (h *participantsHandler) CreateParticipant(w http.ResponseWriter, r *http.Request) {
	participant := &models.Participant{}
	err := json.NewDecoder(r.Body).Decode(participant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.CreateParticipant(participant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *participantsHandler) DeleteParticipant(w http.ResponseWriter, r *http.Request) {
	classroomID, err := strconv.Atoi(r.PathValue("classroomID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteParticipant(classroomID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
