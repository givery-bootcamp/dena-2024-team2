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

func (uc *DeleteChannelPostUsecase) Execute(ctx *gin.Context, serverId int, channelId int, postId int) error {
	serverRepo := repositories.NewGetServerRepository(DB(ctx))
	if _, err := serverRepo.Get(serverId); err != nil {
		return err
	}

	channelRepo := repositories.NewGetChannelRepository(DB(ctx))
	if _, err := channelRepo.Get(channelId); err != nil {
		return err
	}

	postRepo := repositories.NewDeletePostRepository(DB(ctx))
	err := postRepo.Delete(postId)
	if err != nil {
		return err
	}
	return nil
}
