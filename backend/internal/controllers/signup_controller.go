package controllers

import (
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
	usecase.Execute(ctx)

}

// name or password のどちらかが不足のときエラーにする
func validateSignUpParameters(account entities.Account) error {
	if account.Name == "" || account.Password == "" {
		return fmt.Errorf("%v", "Missing name or password")
	}
	return nil
}
