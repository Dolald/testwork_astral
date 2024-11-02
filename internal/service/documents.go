package service

import (
	webСache "web-cache/internal/domain"
	"web-cache/internal/repository"
)

type DocumentService struct {
	repo repository.Document
}

func NewDocumentService(repo repository.Document) *DocumentService {
	return &DocumentService{repo: repo}
}

func (t *DocumentService) CreateDocument(userId int, document webСache.Document) (int, error) {
	return t.repo.CreateDocument(userId, document)
}

func (t *DocumentService) GetAllDocuments(userId int) ([]webСache.Document, error) {
	return t.repo.GetAllDocuments(userId)
}

func (t *DocumentService) GetById(userId, documentId int) (webСache.Document, error) {
	return t.repo.GetById(userId, documentId)
}

func (t *DocumentService) DeleteDocument(userId, documentId int) error {
	return t.repo.DeleteDocument(userId, documentId)
}
