package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"
)

type CreatePostUsecase struct {
	repository interfaces.CreatePostRepository
}

func NewCreatePostUsecase(r interfaces.CreatePostRepository) *CreatePostUsecase {
	return &CreatePostUsecase{
		repository: r,
	}
}

func (u *CreatePostUsecase) Execute(post entities.Post) (*entities.Post, error) {
	return u.repository.Post(post)
}
