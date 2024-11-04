package service

import (
	"context"
	"fmt"

	"github.com/Dolald/testwork_astral/internal/models"
	"github.com/Dolald/testwork_astral/internal/repository"
)

type DocumentService struct {
	repo repository.Document
}

func NewDocumentService(repo repository.Document) *DocumentService {
	return &DocumentService{repo: repo}
}

func (t *DocumentService) CreateDocument(ctx context.Context, userId int, document models.Document) (int, error) {
	id, err := t.repo.CreateDocument(ctx, userId, document)
	if err != nil {
		return 0, fmt.Errorf("CreateDocument: %w", err)
	}

	return id, nil
}

func (t *DocumentService) GetAllDocuments(ctx context.Context, userId int, filteredDocuments models.Filters) ([]models.DocumentsResponse, error) {
	documentsList, err := t.repo.GetAllDocuments(ctx, userId, filteredDocuments)
	if err != nil {
		return nil, fmt.Errorf("GetAllDocuments: %w", err)
	}

	return documentsList, nil
}

func (t *DocumentService) GetDocumentById(ctx context.Context, userId, documentId int) (models.Document, error) {
	document, err := t.repo.GetDocumentById(ctx, userId, documentId)
	if err != nil {
		return models.Document{}, fmt.Errorf("GetDocumentById: %w", err)
	}

	return document, nil
}

func (t *DocumentService) DeleteDocument(ctx context.Context, userId, documentId int) error {
	err := t.repo.DeleteDocument(ctx, userId, documentId)
	if err != nil {
		return fmt.Errorf("DeleteDocument: %w", err)
	}

	return nil
}
