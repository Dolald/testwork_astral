package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Dolald/testwork_astral/internal/models"

	"github.com/jmoiron/sqlx"
)

type DocumentPostgres struct {
	db *sqlx.DB
}

func NewDocumentsPostgres(db *sqlx.DB) *DocumentPostgres {
	return &DocumentPostgres{db: db}
}

func (t *DocumentPostgres) CreateDocument(ctx context.Context, userId int, document models.Document) (int, error) {
	createListQuery := "INSERT INTO documents (user_id, filename, url) VALUES ($1, $2, $3) RETURNING id"
	row := t.db.QueryRowContext(ctx, createListQuery, userId, document.Filename, document.Url)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("id scan failed: %w", err)
	}

	return id, nil
}

func (t *DocumentPostgres) GetAllDocuments(ctx context.Context, userId int, allDocuments models.Filters) ([]models.DocumentsResponse, error) {
	var args []string
	var documents []models.DocumentsResponse

	if allDocuments.SortByDate {
		args = append(args, "created_at ASC")
	}

	if allDocuments.SortByName {
		args = append(args, "filename ASC")
	}

	allDocumentsQuery := "SELECT filename, url, created_at FROM documents WHERE user_id = $1"

	if len(args) > 0 {
		allDocumentsQuery += " ORDER BY " + strings.Join(args, ", ")
	}

	if allDocuments.LimitDocuments > 0 {
		allDocumentsQuery += fmt.Sprintf(" LIMIT %d", allDocuments.LimitDocuments)
	}

	err := t.db.SelectContext(ctx, &documents, allDocumentsQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("select failed: %w", err)
	}

	return documents, nil
}

func (t *DocumentPostgres) GetDocumentById(ctx context.Context, userId, documentId int) (models.Document, error) {
	var list models.Document

	getOneList := "SELECT filename, url FROM documents dt WHERE dt.user_id = $1 AND dt.id = $2"

	if err := t.db.GetContext(ctx, &list, getOneList, userId, documentId); err != nil {
		return models.Document{}, fmt.Errorf("get failed: %w", err)
	}

	if (list == models.Document{}) {
		return list, sql.ErrNoRows
	}

	return list, nil
}

func (t *DocumentPostgres) DeleteDocument(ctx context.Context, userId, documentId int) error {
	query := "DELETE FROM documents dt WHERE $1 = dt.user_id AND dt.id = $2"

	result, err := t.db.ExecContext(ctx, query, userId, documentId)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no documents were changed: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Document not found")
	}

	return nil
}
