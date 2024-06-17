package interfaces

import (
	"myapp/internal/entities"
)

type CreatePostRepository interface {
	Post(post entities.Post) (*entities.Post, error)
}
