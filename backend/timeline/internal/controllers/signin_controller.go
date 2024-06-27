package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signinReqestParam struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Signin(ctx *gin.Context) {
	var json signinReqestParam
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	uc := usecases.NewSigninUsecase()
	token, err := uc.Execute(ctx, json.Name, json.Password)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			handleError(ctx, http.StatusUnauthorized, err)
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
