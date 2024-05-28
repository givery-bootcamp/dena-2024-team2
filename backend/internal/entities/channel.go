package entities

import "time"

type Channels []Channel
type Channel struct {
	Id        int
	ServerId  int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}