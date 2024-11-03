package repository

import (
	"database/sql"
	"fmt"
	"strings"
	webCache "web-cache/internal/domain"

	"github.com/jmoiron/sqlx"
)

type DocumentPostgres struct {
	db *sqlx.DB
}

func NewDocumentsPostgres(db *sqlx.DB) *DocumentPostgres {
	return &DocumentPostgres{db: db}
}

func (t *DocumentPostgres) CreateDocument(userId int, document webCache.Document) (int, error) {

	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, filename, url) VALUES ($1, $2, $3) RETURNING id", documentsTable)
	row := t.db.QueryRow(createListQuery, userId, document.Filename, document.Url)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *DocumentPostgres) GetAllDocuments(userId int, allDocuments webCache.Filters) ([]webCache.DocumentsResponse, error) {
	var args []string
	var documents []webCache.DocumentsResponse

	if allDocuments.SortByDate {
		args = append(args, "created_at ASC")
	}

	if allDocuments.SortByName {
		args = append(args, "filename ASC")
	}

	allDocumentsQuery := fmt.Sprintf("SELECT filename, url, created_at FROM %s WHERE user_id = $1", documentsTable)

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

	fmt.Println(documents)

	return documents, nil
}

func (t *DocumentPostgres) GetById(userId, documentId int) (webCache.Document, error) {
	var list webCache.Document

	getOneList := fmt.Sprintf("SELECT filename, url FROM %s dt WHERE dt.user_id = $1 AND dt.id = $2", documentsTable)

	if err := t.db.Get(&list, getOneList, userId, documentId); err != nil {
		return webCache.Document{}, err
	}

	if (list == webCache.Document{}) {
		return list, sql.ErrNoRows
	}

	return list, nil
}

func (t *DocumentPostgres) DeleteDocument(userId, documentId int) error {
	query := fmt.Sprintf("DELETE FROM %s dt WHERE  $1 = dt.user_id AND dt.id = $2", documentsTable)

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
