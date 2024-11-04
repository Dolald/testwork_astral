package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Dolald/testwork_astral/internal/domain"
	"github.com/jmoiron/sqlx"
)

type DocumentPostgres struct {
	db *sqlx.DB
}

func NewDocumentsPostgres(db *sqlx.DB) *DocumentPostgres {
	return &DocumentPostgres{db: db}
}

func (t *DocumentPostgres) CreateDocument(userId int, document domain.Document) (int, error) {

	createListQuery := "INSERT INTO documents (user_id, filename, url) VALUES ($1, $2, $3) RETURNING id"
	row := t.db.QueryRow(createListQuery, userId, document.Filename, document.Url)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *DocumentPostgres) GetAllDocuments(userId int, allDocuments domain.Filters) ([]domain.DocumentsResponse, error) {
	var args []string
	var documents []domain.DocumentsResponse

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

	err := t.db.Select(&documents, allDocumentsQuery, userId)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (t *DocumentPostgres) GetDocumentById(userId, documentId int) (domain.Document, error) {
	var list domain.Document

	getOneList := "SELECT filename, url FROM documents dt WHERE dt.user_id = $1 AND dt.id = $2"

	if err := t.db.Get(&list, getOneList, userId, documentId); err != nil {
		return domain.Document{}, err
	}

	if (list == domain.Document{}) {
		return list, sql.ErrNoRows
	}

	return list, nil
}

func (t *DocumentPostgres) DeleteDocument(userId, documentId int) error {
	query := "DELETE FROM documents dt WHERE  $1 = dt.user_id AND dt.id = $2"

	result, err := t.db.Exec(query, userId, documentId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Document not found")
	}

	return nil
}
