package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}

func NewUser(id int, firstName, lastName, email, role, password string) *User {
	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Role:      role,
		Password:  password,
	}
}
