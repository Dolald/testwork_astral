package service

import (
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type Authorization interface {
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	CreateUser(user domain.User) (int, error)
}

type Document interface {
	CreateDocument(userId int, document domain.Document) (int, error)
	GetAllDocuments(userId int, filteredDocuments domain.Filters) ([]domain.DocumentsResponse, error)
	GetDocumentById(userId, documentId int) (domain.Document, error)
	DeleteDocument(userId, documentId int) error
}

type Service struct {
	Authorization
	User
	Document
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Document:      NewDocumentService(repos.Document),
		User:          NewUserService(repos.Document),
	}
}
