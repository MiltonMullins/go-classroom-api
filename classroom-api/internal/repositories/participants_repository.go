package repositories

import (
	"database/sql"

	"github.com/miltonmullins/classroom-api/classroom-api/internal/models"
)

type ParticipantsRepository interface {
	GetParticipantsByClassroomID(classroomID int) ([]*models.ParticipantResponse, error)
	CreateParticipant(participant *models.Participant) error
	DeleteParticipant(classroomID, userID int) error
}

type participantsRepository struct {
	db *sql.DB
}

func NewParticipantsRepository(db *sql.DB) ParticipantsRepository {
	return &participantsRepository{
		db: db,
	}
}

func (r *participantsRepository) GetParticipantsByClassroomID(classroomID int) ([]*models.ParticipantResponse, error) {
	rows, err := r.db.Query("SELECT classroom_id, user_id, role FROM participants WHERE classroom_id = $1", classroomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	participants := []*models.ParticipantResponse{}
	for rows.Next() {
		participant := &models.ParticipantResponse{}
		err := rows.Scan(&participant.ClassroomID, &participant.UserID, &participant.Role)
		if err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return participants, nil
}

func (r *participantsRepository) CreateParticipant(participant *models.Participant) error {
	_, err := r.db.Exec("INSERT INTO participants (classroom_id, user_id, role) VALUES ($1, $2, $3)", participant.ClassroomID, participant.UserID, participant.Role)
	if err != nil {
		return err
	}

	return nil
}

func (r *participantsRepository) DeleteParticipant(classroomID, userID int) error {
	_, err := r.db.Exec("DELETE FROM participants WHERE classroom_id = $1 AND user_id = $2", classroomID, userID)
	if err != nil {
		return err
	}

	return nil
}



