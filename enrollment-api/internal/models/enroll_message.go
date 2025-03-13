package models

type EnrollMessage struct {
	AssigmentTitle string `json:"assigment_title"`
	StudentID      string `json:"student_id"`
}

func NewEnrollMessage(assigmentTitle, studentID string) *EnrollMessage {
	return &EnrollMessage{
		AssigmentTitle: assigmentTitle,
		StudentID:      studentID,
	}
}
