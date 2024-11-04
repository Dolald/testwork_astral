package service

import (
	"context"
	"fmt"

	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (a *UserService) CreateUser(ctx context.Context, user models.User) (int, error) {
	user.Password = gereratePasswordHash(user.Password)

	id, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("createUser failed")
	}

	return id, nil
}
