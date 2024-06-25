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

func (uc *CreateServerUsecase) Execute(ctx *gin.Context, name string, icon string) (entities.Server, error) {
	serverRepo := repositories.NewServerRepository()
}
