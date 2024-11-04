package service

import (
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type UserService struct {
	repo repository.User // Предполагается, что у вас есть репозиторий для пользователей
}

func NewUserService(repo repository.User) *DocumentService {
	return &DocumentService{repo: repo}
}

func (a *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = gereratePasswordHash(user.Password)

	return a.repo.CreateUser(user)
}
