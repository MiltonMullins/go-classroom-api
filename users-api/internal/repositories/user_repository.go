package repositories

import (
	"database/sql"
	//"errors"

	"github.com/miltonmullins/classroom-api/users-api/internal/models"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() []*models.User
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUserById(id int) (*models.User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			//TODO dont throw a panic and stop the service
			panic(err)
			//return nil, errors.New("error scanning user")
		}
	}

	return &user, nil
}

func (u *userRepository) GetUsers() []*models.User {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role, &user.Password)
		if err != nil {
			//TODO dont throw a panic and stop the service
			panic(err)
			//return nil, errors.New("error scanning user")
		}
		users = append(users, &user)
	}

	return users
}

func (u *userRepository) CreateUser(user *models.User) (*models.User, error) {
	_, err := u.db.Exec("INSERT INTO users (first_name, last_name, email, role, password) VALUES ($1, $2, $3, $4, $5)", user.FirstName, user.LastName, user.Email, user.Role, user.Password)
	if err != nil {
		//TODO dont throw a panic and stop the service
		panic(err)
	}

	return user, nil
}

func (u *userRepository) UpdateUser(id int, user *models.User) (*models.User, error) {
	_, err := u.db.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3, role = $4 WHERE id = $5", user.FirstName, user.LastName, user.Email, user.Role, id)
	if err != nil {
		//TODO dont throw a panic and stop the service
		panic(err)
	}

	return user, nil
}

func (u *userRepository) DeleteUser(id int) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		//TODO dont throw a panic and stop the service
		panic(err)
	}

	return nil
}
