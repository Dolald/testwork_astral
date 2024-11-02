package repository

import (
	"database/sql"
	"fmt"
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

func (t *DocumentPostgres) GetAllDocuments(userId int) ([]webCache.Document, error) {
	// var lists []webCache.Document

	// getAllListsQuery := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl JOIN %s ul ON (tl.id = ul.list_id) WHERE ul.user_id = $1", documentsTable, usersTable)

	// err := t.db.Select(&lists, getAllListsQuery, userId)

	return []webCache.Document{}, nil
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

func (t *DocumentPostgres) DeleteDocument(userId, listId int) error {
	// query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND $1 = ul.user_id AND ul.list_id = $2" /* todoListsTable usersListsTable*/)

	// _, err := t.db.Exec(query, userId, listId)

	return nil
}
