package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {

	var account entities.Account
	ctx.BindJSON(&account)
	fmt.Println(account)
	if err := validateSignUpParameters(account); err != nil {
		handleError(ctx, 400, err)
		return
	}

	usecase := usecases.NewSignUpUsecase()
	usecase.Execute()

}

// name or password のどちらかが不足のときエラーにする
func validateSignUpParameters(account entities.Account) error {
	if account.Name == "" || account.Password == "" {
		return errors.New(fmt.Sprintf("Missing name or password"))
	}
	return nil
}
