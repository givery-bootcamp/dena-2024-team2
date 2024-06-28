package interfaces

import (
	"myapp/internal/entities"
)

type CreatePostRepository interface {
	Post(channelID int, userID int, userName string, content string) (*entities.Post, error)
}
