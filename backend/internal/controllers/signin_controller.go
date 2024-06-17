package controllers

import (
	"fmt"
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
	user, err := uc.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "login success"})
}
