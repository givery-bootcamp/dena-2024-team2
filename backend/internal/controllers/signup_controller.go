package controllers

import (
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {

	var account entities.Account
	if err := ctx.BindJSON(&account); err != nil {
		handleError(ctx, 400, err)
		return
	}
	fmt.Println(account)
	if err := validateSignUpParameters(account); err != nil {
		handleError(ctx, 400, err)
		return
	}

	usecase := usecases.NewSignUpUsecase()
	err := usecase.Execute(ctx, account.Name, account.Password)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
	} else {
		ctx.JSON(200, "suceess")
	}
}

// name or password のどちらかが不足のときエラーにする
func validateSignUpParameters(account entities.Account) error {
	if account.Name == "" || account.Password == "" {
		return fmt.Errorf("%v", "Missing name or password")
	}
	return nil
}
