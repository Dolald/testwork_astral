package service

import (
	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type DocumentService struct {
	repo repository.Document
}

func NewDocumentService(repo repository.Document) *DocumentService {
	return &DocumentService{repo: repo}
}

func (t *DocumentService) CreateDocument(userId int, document domain.Document) (int, error) {
	return t.repo.CreateDocument(userId, document)
}

func (t *DocumentService) GetAllDocuments(userId int, filteredDocuments domain.Filters) ([]domain.DocumentsResponse, error) {
	return t.repo.GetAllDocuments(userId, filteredDocuments)
}

func (t *DocumentService) GetDocumentById(userId, documentId int) (domain.Document, error) {
	return t.repo.GetDocumentById(userId, documentId)
}

func (t *DocumentService) DeleteDocument(userId, documentId int) error {
	return t.repo.DeleteDocument(userId, documentId)
}
