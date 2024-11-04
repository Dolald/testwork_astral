package repository

import (
	"context"
	"fmt"

	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, user models.User) (int, error) {
	var id int

	query := "INSERT INTO users (login, password_hash) values ($1, $2) RETURNING id"
	row := r.db.QueryRowContext(ctx, query, user.Login, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("scan failed: %w", err)
	}
	return id, nil
}

func (r *UserPostgres) GetUser(ctx context.Context, login, password string) (models.User, error) {
	var user models.User

	query := "SELECT id FROM users WHERE login=$1 AND password_hash=$2"
	err := r.db.GetContext(ctx, &user, query, login, password)
	if err != nil {
		return models.User{}, fmt.Errorf("get failed: %w", err)
	}

	return user, nil
}
