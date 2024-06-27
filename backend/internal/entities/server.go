package entities

import (
	"time"
)

type Servers struct {
	Servers []Server `json:"servers"`
}
type Server struct {
	Id        int       `json:"id"`
	OwnerId   int       `json:"serverId"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
