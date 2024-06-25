package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createChannelsRequestParam struct {
	Name string `json:"name"`
}

func CreateChannels(ctx *gin.Context) {
	var param createChannelsRequestParam
	serverId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.BindJSON(&param); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	// チャンネル名がない場合
	if param.Name == "" {
		handleError(ctx, http.StatusBadRequest, err)
	}

	serverRepository := repositories.NewGetServerRepository(DB(ctx))
	createChannelRepository := repositories.NewCreateChannelRepository(DB(ctx))
	usecase := usecases.NewCreateChannelUsecase(serverRepository, createChannelRepository)
	result, err := usecase.Execute(serverId, param.Name)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
	} else {
		ctx.JSON(200, result)
	}

}
