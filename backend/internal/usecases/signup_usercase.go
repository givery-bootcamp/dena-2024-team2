package usecases

import (
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type SignUpUsecase struct {
}

func NewSignUpUsecase() *SignUpUsecase {
	return &SignUpUsecase{}
}

func (uc *SignUpUsecase) Execute(ctx *gin.Context, name string, password string) error {
	repository := repositories.NewSignUpRepository(DB(ctx))
	err := repository.SignUp(name, password)

	return err
}
