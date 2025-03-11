package models

type Participant struct {
	ClassroomID int
	UserID      int
	Role		string
}

func NewParticipant(classroomID, userID int, role string) *Participant {
	return &Participant{
		ClassroomID: classroomID,
		UserID:      userID,
		Role:		 role,
	}
}