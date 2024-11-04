package service

import (
	"fmt"

	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (a *UserService) CreateUser(user domain.User) (int, error) {
	user.Password = gereratePasswordHash(user.Password)

	id, err := a.repo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("createUser failed")
	}

	return id, nil
}
