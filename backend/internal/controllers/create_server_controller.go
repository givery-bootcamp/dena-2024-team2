package controllers

import (
	"net/http"

	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

type createServerRequestParams struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func CreateServer(ctx *gin.Context) {
	var req createServerRequestParams
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	uid := getUserIdFromContext(ctx)
	uc := usecases.NewCreateServerUsecase()
	server, err := uc.Execute(ctx, req.Name, req.Icon, uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, server)
}
