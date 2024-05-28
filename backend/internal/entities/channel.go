package entities

import (
	"time"
)

type Channel struct {
	Id        int
	ServerId  int
	Name      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
