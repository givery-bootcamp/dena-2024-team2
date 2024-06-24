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

func (uc *GetChannelPostUsecase) Execute(ctx *gin.Context, channelId int) (entities.Posts, error) {
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
