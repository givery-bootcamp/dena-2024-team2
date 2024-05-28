package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myapp/internal/repositories"
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

