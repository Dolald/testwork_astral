package repository

import (
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Document interface {
	CreateDocument(userId int, document domain.Document) (int, error)
	GetAllDocuments(userId int, filteredDocuments domain.Filters) ([]domain.DocumentsResponse, error)
	GetDocumentById(userId, documentId int) (domain.Document, error)
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
