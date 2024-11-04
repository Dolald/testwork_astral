package repository

import (
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Document interface {
	CreateDocument(userId int, document domain.Document) (int, error)
	GetAllDocuments(userId int, filteredDocuments models.Filters) ([]models.DocumentsResponse, error)
	GetDocumentById(userId, documentId int) (domain.Document, error)
	DeleteDocument(userId, documentId int) error
}

type Repository struct {
	Document
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Document: NewDocumentsPostgres(db),
		User:     NewUserPostgres(db),
	}
}
