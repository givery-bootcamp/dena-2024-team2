package usecases

import (
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type SignupUsecase struct {
}

func NewSignupUsecase() *SignupUsecase {
	return &SignupUsecase{}
}

func (uc *SignupUsecase) Execute(ctx *gin.Context, name string, password string) error {
	repository := repositories.NewSignupRepository(DB(ctx))
	err := repository.Signup(name, password)

	return err
}
