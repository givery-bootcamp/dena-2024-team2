package usecases

import "github.com/gin-gonic/gin"

type CreateChannelUsecase struct {
}

func NewCreateChannelUsecase() *CreateChannelUsecase {
	return &CreateChannelUsecase{}
}

func (uc *CreateChannelUsecase) Execute(ctx *gin.Context, serverId int) error {
	return nil
}
