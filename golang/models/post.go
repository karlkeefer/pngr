package models

import (
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	AuthorID  int64     `json:"author_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Status int

const (
	StatusPrivate Status = 0
	StatusPublic         = 1
)
