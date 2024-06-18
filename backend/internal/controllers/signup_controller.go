package controllers

import (
	"fmt"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Signup(ctx *gin.Context) {

	var account requestParam
	if err := ctx.BindJSON(&account); err != nil {
		handleError(ctx, 400, err)
		return
	}
	fmt.Println(account)
	if err := validateSignupParameters(account); err != nil {
		handleError(ctx, 400, err)
		return
	}

	usecase := usecases.NewSignupUsecase()
	err := usecase.Execute(ctx, account.Name, account.Password)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
	} else {
		ctx.JSON(200, "suceess")
	}
}

// name or password のどちらかが不足のときエラーにする
func validateSignupParameters(account requestParam) error {
	if account.Name == "" || account.Password == "" {
		return fmt.Errorf("%v", "Missing name or password")
	}
	return nil
}
