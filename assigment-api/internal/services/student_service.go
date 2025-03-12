package services

import (
	"context"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/models"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/repositories"
)

type StudentTaskService interface {
	GetStudentTasks(ctx context.Context, assigmentID int) ([]*models.StudentTask, error)
	CreateStudentTask(ctx context.Context, studentTask *models.StudentTask) error
	UpdateStudentTask(ctx context.Context, studentID, assigmentID int, studentTask *models.StudentTask) error
	DeleteStudentTask(ctx context.Context, studentID, assigmentID int) error
}

type studentTaskService struct {
	studentTaskRepository repositories.StudentTaskRepository
}

func NerStudentTaskRepository(studentTaskRepository repositories.StudentTaskRepository) StudentTaskService {
	return &studentTaskService {
		studentTaskRepository: studentTaskRepository,
	}
}

func (s *studentTaskService) GetStudentTasks(ctx context.Context, assigmentID int) ([]*models.StudentTask, error) {
	return s.studentTaskRepository.GetStudentTasks(ctx, assigmentID)
}
func (s *studentTaskService) CreateStudentTask(ctx context.Context, studentTask *models.StudentTask) error {
	return s.studentTaskRepository.CreateStudentTask(ctx, studentTask)
}
func (s *studentTaskService) UpdateStudentTask(ctx context.Context, studentID, assigmentID int, studentTask *models.StudentTask) error {
	return s.studentTaskRepository.UpdateStudentTask(ctx, studentID, assigmentID, studentTask)
}
func (s *studentTaskService) DeleteStudentTask(ctx context.Context, studentID, assigmentID int) error {
	return s.studentTaskRepository.DeleteStudentTask(ctx, studentID, assigmentID)
}