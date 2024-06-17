package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/entities"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {

}

// name or password のどちらかが不足のときエラーにする
func validateSignUpParameters(user entities.User) error {
	if user.Name == "" || user.Password == "" {
		return errors.New(fmt.Sprintf("Missing name or password"))
	}
	return nil
}
