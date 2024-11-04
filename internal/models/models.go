package models

import "time"

type User struct {
	Id       int    `json:"-" db:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Document struct {
	Id         int       `json:"id" db:"id"`
	User_id    int       `json:"user_id" db:"user_id"`
	Filename   string    `json:"filename" db:"filename"`
	Url        string    `json:"url" db:"url"`
	Created_at time.Time `json:"-" db:"created_at"`
}

type DocumentsResponse struct {
	Filename   string    `json:"filename" db:"filename"`
	Url        string    `json:"url" db:"url"`
	Created_at time.Time `json:"created_at" db:"created_at"`
}

type Filters struct {
	LimitDocuments int  `json: "limitDocuments"`
	SortByName     bool `json: "sortByName"`
	SortByDate     bool `json: "sortByDate"`
}
