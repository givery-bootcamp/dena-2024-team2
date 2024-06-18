package controllers

import (
	"errors"
	"myapp/internal/entities"
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
	uc := usecases.NewSigninUsecase()
	token, err := uc.Execute(ctx, json.UserName, json.Password)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			handleError(ctx, http.StatusUnauthorized, errors.New("failed to signin"))
		} else {
			handleError(ctx, http.StatusInternalServerError, err)
		}
		return
	}

	ctx.JSON(http.StatusOK, requestJson{token})
}

type requestJson struct {
	AccessToken string `json:"access_token"`
}
