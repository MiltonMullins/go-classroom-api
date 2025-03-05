package services

import (
	"github.com/miltonmullins/classroom-api/users-api/internal/models"
	"github.com/miltonmullins/classroom-api/users-api/internal/repositories"
)

type UserService interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() []*models.User
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id int, user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) GetUserById(id int) (*models.User, error) {
	return u.userRepository.GetUserById(id)
}

func (u *userService) GetUsers() []*models.User {
	return u.userRepository.GetUsers()
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	return u.userRepository.CreateUser(user)
}

func (u *userService) UpdateUser(id int, user *models.User) (*models.User, error) {
	return u.userRepository.UpdateUser(id, user)
}

func (u *userService) DeleteUser(id int) error {
	return u.userRepository.DeleteUser(id)
}
