package interfaces

import "myapp/internal/entities"

type PostRepository interface {
	Get(postId int) (*entities.Post, error)
	Update(postId int, content string) (*entities.Post, error)
}
