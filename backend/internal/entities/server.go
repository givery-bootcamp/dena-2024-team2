package entities

import "time"

type Server struct {
	Id        int
	OwnerId   int
	Name      string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
