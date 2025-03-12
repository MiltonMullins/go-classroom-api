package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/models"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/services"
)

type StudentTaskHandler interface {
	GetStudentTasks(w http.ResponseWriter, r *http.Request) 
	CreateStudentTask(w http.ResponseWriter, r *http.Request) 
	UpdateStudentTask(w http.ResponseWriter, r *http.Request) 
	DeleteStudentTask(w http.ResponseWriter, r *http.Request) 
}

type studenTaskHandler struct {
	studentTaskService services.StudentTaskService
}

func NewStudentTasklHandler(studentTaskService services.StudentTaskService) StudentTaskHandler {
	return &studenTaskHandler{
		studentTaskService: studentTaskService,
	}
}

func (h *studenTaskHandler) GetStudentTasks(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	assigmentID, err := strconv.Atoi(r.PathValue("assigment_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	studenStasks, err:= h.studentTaskService.GetStudentTasks(context, assigmentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonStudentTask, err := json.Marshal(studenStasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonStudentTask)
}
func (h *studenTaskHandler) CreateStudentTask(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	var studenStasks models.StudentTask
	err := json.NewDecoder(r.Body).Decode(&studenStasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.studentTaskService.CreateStudentTask(context, &studenStasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *studenTaskHandler) UpdateStudentTask(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	studentID, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	assigmentID, err := strconv.Atoi(r.PathValue("assigment_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var studentTask models.StudentTask
	err = json.NewDecoder(r.Body).Decode(&studentTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.studentTaskService.UpdateStudentTask(context, studentID, assigmentID, &studentTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (h *studenTaskHandler) DeleteStudentTask(w http.ResponseWriter, r *http.Request) {
	context := context.Background()
	studentID, err := strconv.Atoi(r.PathValue("student_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	assigmentID, err := strconv.Atoi(r.PathValue("assigment_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.studentTaskService.DeleteStudentTask(context, studentID, assigmentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}