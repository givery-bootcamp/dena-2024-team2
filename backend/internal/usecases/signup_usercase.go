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

func (uc *SignUpUsecase) Execute(ctx *gin.Context) error {
	repository := repositories.NewSignUpRepository(DB(ctx))
	repository.SignUp()

	return nil
}
