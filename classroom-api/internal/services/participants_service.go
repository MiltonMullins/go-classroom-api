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
	participantsRepository repositories.ParticipantsRepository
}

func NewParticipantsService(repo repositories.ParticipantsRepository) ParticipantsService {
	return &participantsService{
		participantsRepository: repo,
	}
}

func (s *participantsService) GetParticipantsByClassroomID(classroomID int) ([]*models.Participant, error) {
	return s.participantsRepository.GetParticipantsByClassroomID(classroomID)
}

func (s *participantsService) CreateParticipant(participant *models.Participant) error {
	return s.participantsRepository.CreateParticipant(participant)
}

func (s *participantsService) DeleteParticipant(classroomID, userID int) error {
	return s.participantsRepository.DeleteParticipant(classroomID, userID)
}