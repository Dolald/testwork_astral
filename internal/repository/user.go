package repository

import (
	"fmt"

	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user domain.User) (int, error) {
	var id int

	query := "INSERT INTO users (login, password_hash) values ($1, $2) RETURNING id"
	row := r.db.QueryRow(query, user.Login, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("scan failed")
	}
	return id, nil
}

func (r *UserPostgres) GetUser(login, password string) (domain.User, error) {
	var user domain.User

	query := "SELECT id FROM users WHERE login=$1 AND password_hash=$2"
	err := r.db.Get(&user, query, login, password)
	if err != nil {
		return domain.User{}, fmt.Errorf("get failed")
	}

	return user, nil
}
