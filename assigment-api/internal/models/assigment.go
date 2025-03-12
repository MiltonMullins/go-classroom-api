package models

import "time"

type Assigment struct {
	Title     string    `json:"title" bson:"title,omitempty"`
	DueDate   time.Time `json:"due_date" bson:"due_date,omitempty"`
	Mandatory bool      `json:"mandatory" bson:"mandatory,omitempty"`
}

func NewAssigment(title string, dueDate time.Time, mandatory bool) *Assigment {
	return &Assigment{
		Title:     title,
		DueDate:   dueDate,
		Mandatory: mandatory,
	}
}
