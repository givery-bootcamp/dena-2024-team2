package controllers

import (
	"fmt"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signupRequestParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Signup(ctx *gin.Context) {

	var account signupRequestParam
	if err := ctx.BindJSON(&account); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := validateSignupParameters(account); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	usecase := usecases.NewSignupUsecase()
	err := usecase.Execute(ctx, account.Name, account.Password)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
	} else {
		ctx.JSON(200, "アカウント登録完了")
	}
}

// name or password のどちらかが不足のときエラーにする
func validateSignupParameters(account signupRequestParam) error {
	if account.Name == "" || account.Password == "" {
		return fmt.Errorf("%v", "名前かパスワードが不足しています。")
	}
	return nil
}
