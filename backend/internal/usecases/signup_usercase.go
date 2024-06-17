package usecases

import (
	"github.com/gin-gonic/gin"
)

type SignUpUsecase struct {
}

func NewSignUpUsecase() *SignUpUsecase {
	return &SignUpUsecase{}
}

func (uc *SignUpUsecase) Execute(ctx *gin.Context) error {

	return nil
}
