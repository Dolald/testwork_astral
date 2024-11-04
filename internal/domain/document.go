package domain

import "time"

type Document struct {
	Id         int
	User_id    int
	Filename   string
	Url        string
	Created_at time.Time
}
