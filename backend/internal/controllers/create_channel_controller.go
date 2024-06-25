package controllers

import (
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

	usecase := usecases.NewCreateChannelUsecase()
	result, err := usecase.Execute(ctx, serverId, param.Name)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
	} else {
		ctx.JSON(200, result)
	}

}
