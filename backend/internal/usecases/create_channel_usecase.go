package usecases

import (
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type CreateChannelUsecase struct {
}

func NewCreateChannelUsecase() *CreateChannelUsecase {
	return &CreateChannelUsecase{}
}

func (uc *CreateChannelUsecase) Execute(ctx *gin.Context, serverId int, name string) error {

	repository := repositories.NewCreateChannelRepository(DB(ctx))
	err := repository.Post(serverId, name)
	return err
}
