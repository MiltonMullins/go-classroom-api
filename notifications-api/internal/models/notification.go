package models

type Notification struct {
	UserID         int    `json:"user_id"`
	AssigmentTitle string `json:"assigment_title"`
	Message        string `json:"message"`
}
