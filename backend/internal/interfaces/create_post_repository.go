package interfaces

import (
	"myapp/internal/entities"
)

type CreatePostRepository interface {
	Post(userID int, channelID int, content string) (*entities.Post, error)
}
