package controllers

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/repositories"
)

func GetChannels(ctx *gin.Context) {
	repository := repositories.NewChannelsRepository(DB(ctx))
	repository.Get()
    ctx.JSON(200, "")
}

