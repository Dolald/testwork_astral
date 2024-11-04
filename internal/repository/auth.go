package repository

import (
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int

	query := "INSERT INTO users (login, password_hash) values ($1, $2) RETURNING id"
	row := r.db.QueryRow(query, user.Login, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (domain.User, error) {
	var user domain.User

	query := "SELECT id FROM users WHERE login=$1 AND password_hash=$2"
	err := r.db.Get(&user, query, login, password)

	return user, err
}
