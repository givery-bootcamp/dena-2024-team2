package entities

import (
	"time"
)

type Posts []Post
type Post struct {
	Id        int
	Title     string
	Body      string
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
