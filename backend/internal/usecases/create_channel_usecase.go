package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type CreateChannelUsecase struct {
}

func NewCreateChannelUsecase() *CreateChannelUsecase {
	return &CreateChannelUsecase{}
}

func (uc *CreateChannelUsecase) Execute(ctx *gin.Context, serverId int, name string) (*entities.Channel, error) {
	// 存在しているサーバーか確認する
	serverRepository := repositories.NewGetServerRepository(DB(ctx))
	if _, err := serverRepository.Get(serverId); err != nil {
		return nil, err
	}

	repository := repositories.NewCreateChannelRepository(DB(ctx))
	result, err := repository.Post(serverId, name)
	return result, err
}
