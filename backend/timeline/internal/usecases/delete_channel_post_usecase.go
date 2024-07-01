package usecases

import (
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type DeleteChannelPostUsecase struct {
}

func NewDeleteChannelPostUsecase() *DeleteChannelPostUsecase {
	return &DeleteChannelPostUsecase{}
}

func (uc *DeleteChannelPostUsecase) Execute(ctx *gin.Context, postId int) error {
	postRepo := repositories.NewDeletePostRepository(DB(ctx))
	err := postRepo.Delete(postId)
	if err != nil {
		return err
	}
	return nil
}
