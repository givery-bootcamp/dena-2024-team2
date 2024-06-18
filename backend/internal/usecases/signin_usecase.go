package usecases

import (
	"myapp/internal/infrastructures"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type SigninUsecase struct{}

func NewSigninUsecase() *SigninUsecase {
	return &SigninUsecase{}
}

func (uc *SigninUsecase) Execute(ctx *gin.Context, userName string, password string) (string, error) {
	authRepo := repositories.NewUserRepository(DB(ctx))
	userId, err := authRepo.Authenticate(userName, password)
	if err != nil {
		return "", err
	}
	token, err := infrastructures.GenerateToken(userId)
	if err != nil {
		return "", err
	}
	return token, nil
}
