package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/entities"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {

	var user entities.User
	ctx.BindJSON(&user)
	fmt.Println(user)
	if err := validateSignUpParameters(user); err != nil {
		handleError(ctx, 400, err)
		return
	}

}

// name or password のどちらかが不足のときエラーにする
func validateSignUpParameters(user entities.User) error {
	if user.Name == "" || user.Password == "" {
		return errors.New(fmt.Sprintf("Missing name or password"))
	}
	return nil
}
