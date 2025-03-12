package services

import (
	"context"

	"github.com/miltonmullins/classroom-api/assigment-api/internal/models"
	"github.com/miltonmullins/classroom-api/assigment-api/internal/repositories"
)

type AssigmentService interface {
	GetAssigment(ctx context.Context, param string) ([]*models.Assigment, error)
	CreateAssigment(ctx context.Context, assigment *models.Assigment) error
	UpdateAssigment(ctx context.Context, assigmentID int, assigment *models.Assigment) error
	DeleteAssigment(ctx context.Context, assigmentID int) error
}

type assigmentService struct {
	assigmentRepository repositories.AssigmentRepository
}

func NewAssigmentService(assigmentRepository repositories.AssigmentRepository) AssigmentService {
	return &assigmentService{
		assigmentRepository: assigmentRepository,
	}
}

func (s *assigmentService) GetAssigment(ctx context.Context, param string) ([]*models.Assigment, error) {
	return s.assigmentRepository.GetAssigment(ctx, param)
}
func (s *assigmentService) CreateAssigment(ctx context.Context, assigment *models.Assigment) error {
	return s.assigmentRepository.CreateAssigment(ctx, assigment)
}
func (s *assigmentService) UpdateAssigment(ctx context.Context, assigmentID int, assigment *models.Assigment) error {
	return s.assigmentRepository.UpdateAssigment(ctx, assigmentID, assigment)
}
func (s *assigmentService) DeleteAssigment(ctx context.Context, assigmentID int) error {
	return s.assigmentRepository.DeleteAssigment(ctx, assigmentID)
}
