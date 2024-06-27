package controllers

import (
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"
	"time"

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
		ctx.JSON(200, newChannelConvertToJson(result))
	}

}

func newChannelConvertToJson(v *entities.Channel) newChannelJson {
	return newChannelJson{
		Id:        v.Id,
		ServerId:  v.ServerId,
		Name:      v.Name,
		CreatedAt: v.CreatedAt,
	}
}

type newChannelJson struct {
	Id        int       `json:"id"`
	ServerId  int       `json:"serverId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
