package controllers

import (
	"fmt"
	"net/http"

	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type createServerRequestParams struct {
	name string
	icon string
}

func CreateServer(ctx *gin.Context) {
	var req createServerRequestParams
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	uid, ok := ctx.Get(middleware.userIdKey)
	if !ok {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("failed to retrieve uid from ctx"))
		return
	}
	uc := usecases.NewCreateServerUsecase()
	server, err := uc.Execute(ctx, req.name, req.icon, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, server)
}
