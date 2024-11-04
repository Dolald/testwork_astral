package service

import (
	"fmt"

	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type DocumentService struct {
	repo repository.Document
}

func NewDocumentService(repo repository.Document) *DocumentService {
	return &DocumentService{repo: repo}
}

func (t *DocumentService) CreateDocument(userId int, document domain.Document) (int, error) {
	id, err := t.repo.CreateDocument(userId, document)
	if err != nil {
		return 0, fmt.Errorf("CreateDocument: %w", err)
	}

	return id, nil
}

func (t *DocumentService) GetAllDocuments(userId int, filteredDocuments models.Filters) ([]models.DocumentsResponse, error) {
	documentsList, err := t.repo.GetAllDocuments(userId, filteredDocuments)
	if err != nil {
		return nil, fmt.Errorf("GetAllDocuments: %w", err)
	}

	return documentsList, nil
}

func (t *DocumentService) GetDocumentById(userId, documentId int) (domain.Document, error) {
	document, err := t.repo.GetDocumentById(userId, documentId)
	if err != nil {
		return domain.Document{}, fmt.Errorf("GetDocumentById: %w", err)
	}

	return document, nil
}

func (t *DocumentService) DeleteDocument(userId, documentId int) error {
	err := t.repo.DeleteDocument(userId, documentId)
	if err != nil {
		return fmt.Errorf("DeleteDocument: %w", err)
	}

	return nil
}
