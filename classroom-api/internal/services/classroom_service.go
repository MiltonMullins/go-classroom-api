package services

import (
	"github.com/miltonmullins/classroom-api/classroom-api/internal/models"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/repositories"
)

type ClassroomService interface {
	GetClassrooms() ([]*models.Classroom, error)
	GetClassroomByID(classroomID int) (*models.Classroom, error)
	CreateClassroom(classroom *models.Classroom) error
	UpdateClassroom(classroomID int, classroom *models.Classroom) error
	DeleteClassroom(classroomID int) error
}

type classroomService struct {
	classroomRepository repositories.ClassroomRepository
}

func NewClassroomService(repo repositories.ClassroomRepository) ClassroomService {
	return &classroomService{
		classroomRepository: repo,
	}
}

func (s *classroomService) GetClassrooms() ([]*models.Classroom, error) {
	return s.classroomRepository.GetClassrooms()
}

func (s *classroomService) GetClassroomByID(classroomID int) (*models.Classroom, error) {
	return s.classroomRepository.GetClassroomByID(classroomID)
}

func (s *classroomService) CreateClassroom(classroom *models.Classroom) error {
	return s.classroomRepository.CreateClassroom(classroom)
}

func (s *classroomService) UpdateClassroom(classroomID int, classroom *models.Classroom) error {
	return s.classroomRepository.UpdateClassroom(classroomID, classroom)
}

func (s *classroomService) DeleteClassroom(classroomID int) error {
	return s.classroomRepository.DeleteClassroom(classroomID)
}