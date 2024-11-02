package repository

import (
	webCache "web-cache/internal/domain"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user webCache.User) (int, error)
	GetUser(username, password string) (webCache.User, error)
}

type Document interface {
	CreateDocument(userId int, document webCache.Document) (int, error)
	GetAllDocuments(userId int) ([]webCache.Document, error)
	GetById(userId, documentId int) (webCache.Document, error)
	DeleteDocument(userId, documentId int) error
}

type Repository struct {
	Authorization
	Document
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Document:      NewDocumentsPostgres(db),
	}
}
