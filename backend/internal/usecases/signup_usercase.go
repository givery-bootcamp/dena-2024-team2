package usecases

import (
	"myapp/internal/infrastructures"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

type SignupUsecase struct {
}

func NewSignupUsecase() *SignupUsecase {
	return &SignupUsecase{}
}

func (uc *SignupUsecase) Execute(ctx *gin.Context, name string, password string) error {

	encryptPw, encryptErr := infrastructures.EncryptPassword(password)
	if encryptErr != nil {
		return encryptErr
	}

	repository := repositories.NewSignupRepository(DB(ctx))
	err := repository.Signup(name, encryptPw)

	return err
}
