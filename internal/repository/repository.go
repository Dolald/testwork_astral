package repository

import (
	"context"

	"github.com/Dolald/testwork_astral/internal/models"

	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, username, password string) (models.User, error)
}

type Document interface {
	CreateDocument(ctx context.Context, userId int, document models.Document) (int, error)
	GetAllDocuments(ctx context.Context, userId int, filteredDocuments models.Filters) ([]models.DocumentsResponse, error)
	GetDocumentById(ctx context.Context, userId, documentId int) (models.Document, error)
	DeleteDocument(ctx context.Context, userId, documentId int) error
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
