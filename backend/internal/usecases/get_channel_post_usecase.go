package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type GetChannelPostUsecase struct {
}

func NewGetChannelPostUsecase() *GetChannelPostUsecase {
	return &GetChannelPostUsecase{}
}

func (uc *GetChannelPostUsecase) Execute(ctx *gin.Context, serverId int, channelId int) (entities.Posts, error) {
	serverRepo := repositories.NewGetServerRepository(DB(ctx))
	if _, err := serverRepo.Get(serverId); err != nil {
		return nil, err
	}

	channelRepo := repositories.NewGetChannelRepository(DB(ctx))
	if _, err := channelRepo.Get(channelId); err != nil {
		return nil, err
	}

	postRepo := repositories.NewGetPostRepository(DB(ctx))
	posts, err := postRepo.Get(channelId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
