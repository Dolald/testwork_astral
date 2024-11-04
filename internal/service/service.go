package service

import (
	"context"

	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type Authorization interface {
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
}

type Document interface {
	CreateDocument(ctx context.Context, userId int, document models.Document) (int, error)
	GetAllDocuments(ctx context.Context, userId int, filteredDocuments models.Filters) ([]models.DocumentsResponse, error)
	GetDocumentById(ctx context.Context, userId, documentId int) (models.Document, error)
	DeleteDocument(ctx context.Context, userId, documentId int) error
}

type Service struct {
	Authorization
	User
	Document
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.User),
		Document:      NewDocumentService(repos.Document),
		User:          NewUserService(repos.User),
	}
}
