package entities
import (
	"time"
)

type Posts []Post
type Post struct {
	Id  int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	UserId int `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
