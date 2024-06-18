package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/infrastructures"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func Signin(ctx *gin.Context) {
	var json requestParam
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("%v", json)
	uc := usecases.NewSigninUsecase()
	userId, err := uc.Execute(ctx, json.UserName, json.Password)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			handleError(ctx, http.StatusUnauthorized, errors.New("failed to signin"))
		} else {
			handleError(ctx, http.StatusInternalServerError, err)
		}
		return
	}
	token, err := infrastructures.GenerateToken(userId)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.SetCookie("token", token, 0, "", ctx.Request.Host, false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "login success"})
}
