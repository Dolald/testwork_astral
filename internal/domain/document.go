package domain

import "time"

type Document struct {
	Id         int       `json:"id" db:"id"`
	User_id    int       `json:"id" db:"user_id"`
	Filename   string    `json:"filename" db:"filename"`
	Url        string    `json:"url" db:"url"`
	Created_at time.Time `json:"-" db:"created_at"`
}
