package controllers

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/repositories"
)

func GetPosts(ctx *gin.Context) {
	repository := repositories.NewGetPostRepository(DB(ctx))
	repository.Get()
    ctx.JSON(200, "")
}
