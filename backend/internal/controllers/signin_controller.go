package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func Signin(ctx *gin.Context) {
	var json requestParam
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
