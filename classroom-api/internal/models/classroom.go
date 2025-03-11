package models

type Classroom struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   int    `json:"teacher_id"`
}

func NewClassroom(id int, name, description string, teacherID int) *Classroom {
	return &Classroom{
		ID:          id,
		Name:        name,
		Description: description,
		TeacherID:   teacherID,
	}
}
