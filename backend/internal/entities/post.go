package entities

import "time"

type Post struct {
	Id        int        `json:"id"`
	ChannelId int        `json:"channel_id"`
	UserId    int        `json:"user_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
