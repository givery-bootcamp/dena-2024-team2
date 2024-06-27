package entities

import "time"

type Posts []Post
type Post struct {
	Id        int `json:"id"`
	ChannelId int `json:"channel_id"`
	User      User
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (p *Post) CanEdit(uid int) bool {
	return p.User.Id == uid
}
