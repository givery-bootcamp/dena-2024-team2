package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetChannels(ctx *gin.Context) {
	serverId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	serverRepository := repositories.NewGetServerRepository(DB(ctx))
	getChannelsRepository := repositories.NewGetChannelsRepository(DB(ctx))
	usecase := usecases.NewGetChannelsUsecase()
	channels, err := usecase.Execute()
	if err != nil {
		handleError(ctx, 500, err)
	} else if channels != nil {
		ctx.JSON(200, channelsConvertToJson(channels))
	} else {
		handleError(ctx, 404, errors.New("Not found"))
	}
}

func channelsConvertToJson(channels entities.Channels) ChannelsResponseJson {
	jsonChannels := make([]ChannelJson, len(channels))
	for i, v := range channels {
		jsonChannels[i] = ChannelJson{
			Id:        v.Id,
			ServerId:  v.ServerId,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return ChannelsResponseJson{Channels: jsonChannels}
}

type ChannelsResponseJson struct {
	Channels []ChannelJson `json:"channels"`
}

type ChannelJson struct {
	Id        int       `json:"id"`
	ServerId  int       `json:"serverId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
