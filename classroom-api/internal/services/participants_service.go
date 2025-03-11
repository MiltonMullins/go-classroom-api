package services

import (
	"github.com/miltonmullins/classroom-api/classroom-api/internal/models"
	"github.com/miltonmullins/classroom-api/classroom-api/internal/repositories"
)

type ParticipantsService interface {
	GetParticipantsByClassroomID(classroomID int) ([]*models.Participant, error)
	CreateParticipant(participant *models.Participant) error
	DeleteParticipant(classroomID, userID int) error
}

type participantsService struct {
	repo repositories.ParticipantsRepository
}

func NewParticipantsService(repo repositories.ParticipantsRepository) ParticipantsService {
	return &participantsService{
		repo: repo,
	}
}

func (s *participantsService) GetParticipantsByClassroomID(classroomID int) ([]*models.Participant, error) {
	return s.repo.GetParticipantsByClassroomID(classroomID)
}

func (s *participantsService) CreateParticipant(participant *models.Participant) error {
	return s.repo.CreateParticipant(participant)
}

func (s *participantsService) DeleteParticipant(classroomID, userID int) error {
	return s.repo.DeleteParticipant(classroomID, userID)
}