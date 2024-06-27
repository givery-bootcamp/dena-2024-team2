package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/interfaces"

	"github.com/gin-gonic/gin"
)

type UpdatePostUsecase struct {
	pr interfaces.PostRepository
}

func NewUpdatePostUsecase(pr interfaces.PostRepository) *UpdatePostUsecase {
	return &UpdatePostUsecase{pr: pr}
}

func (uc *UpdatePostUsecase) Execute(ctx *gin.Context, postId int, userId int, content string) (*entities.Post, error) {
	post, err := uc.pr.Get(postId)
	if err != nil {
		return nil, err
	}
	if !post.CanEdit(userId) {
		return nil, err
	}
	post, err = uc.pr.Update(postId, content)
	if err != nil {
		return nil, err
	}
	return post, nil
}
