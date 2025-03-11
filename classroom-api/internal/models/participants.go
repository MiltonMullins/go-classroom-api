package models

type Participant struct {
	ID          int    `json:"id"`
	ClassroomID int    `json:"classroom_id"`
	UserID      int    `json:"user_id"`
	Role        string `json:"role"`
}

func NewParticipant(classroomID, userID int, role string) *Participant {
	return &Participant{
		ClassroomID: classroomID,
		UserID:      userID,
		Role:        role,
	}
}

type ParticipantResponse struct {
	ClassroomID int    `json:"classroom_id"`
	UserID      int    `json:"user_id"`
	Role        string `json:"role"`
}
