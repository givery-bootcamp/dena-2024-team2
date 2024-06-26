package entities

import "time"

type Server struct {
	Id        int       `json:"id"`
	OwnerId   int       `json:"owner_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
