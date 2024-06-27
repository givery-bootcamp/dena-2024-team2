package usecases

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type CreateServerUsecase struct{}

func NewCreateServerUsecase() *CreateServerUsecase {
	return &CreateServerUsecase{}
}

func (uc *CreateServerUsecase) Execute(ctx *gin.Context, name string, icon string, uid int) (*entities.Server, error) {
	serverRepo := repositories.NewGetServerRepository(DB(ctx))
	server, err := serverRepo.Create(ctx, name, icon, uid)
	if err != nil {
		return nil, err
	}
	return server, nil
}
