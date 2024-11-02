package service

import (
	webСache "web-cache/internal/domain"
	"web-cache/internal/repository"
)

type Authorization interface {
	CreateUser(user webСache.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Document interface {
	CreateDocument(userId int, document webСache.Document) (int, error)
	GetAllDocuments(userId int) ([]webСache.Document, error)
	GetById(userId, documentId int) (webСache.Document, error)
	DeleteDocument(userId, documentId int) error
}

type Service struct {
	Authorization
	Document
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Document:      NewDocumentService(repos.Document),
	}
}
