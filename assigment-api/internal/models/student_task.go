package models

type StudentTask struct {
	StudentID   int    `json:"student_id" bson:"student_id,omitempty"`
	AssigmentID int    `json:"assigment_id" bson:"assigment_id,omitempty"`
	Status      string `json:"status" bson:"status,omitempty"`
}

func NewStudentTask(studentID, assigmentID int, status string) *StudentTask {
	return &StudentTask{
		StudentID:   studentID,
		AssigmentID: assigmentID,
		Status:      status,
	}
}
