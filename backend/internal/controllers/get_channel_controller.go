package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myapp/internal/repositories"
	"time"
)

func GetChannels(ctx *gin.Context) {
	repository := repositories.NewChannelsRepository(DB(ctx))
	channels, err := repository.Get()
	if err != nil {
		handleError(ctx, 500, err)
	} else if channels != nil{
		ctx.JSON(200, channels)
	} else {
		handleError(ctx, 404, errors.New("Not found"))
	}
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
	DeletedAt time.Time `json:"deletedAt"`
}
