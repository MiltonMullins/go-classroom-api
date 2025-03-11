package services

import (
	"github.com/miltonmullins/classroom-api/users-api/internal/models"
	"github.com/miltonmullins/classroom-api/users-api/internal/repositories"

	jwt "github.com/miltonmullins/classroom-api/users-api/utils"
)

type LoginService interface {
	Login(user *models.Login) string
}

type loginService struct {
	userRepository repositories.UserRepository
}

func NewLoginService(userRepository repositories.UserRepository) LoginService {
	return &loginService{
		userRepository: userRepository,
	}
}

func (l *loginService) Login(user *models.Login) string {
	users := l.userRepository.GetUsers()

	for _, u := range users {
		if u.Email == user.Email && u.Password == user.Password {
			tokenString, err := jwt.CreateToken(user.Email)
			if err != nil {
				//TODO Handle Error
				return "User or Password are incorrect"
			}
			return tokenString
		}
	}
	return "User or Password are incorrect"
}
