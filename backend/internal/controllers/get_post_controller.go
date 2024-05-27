package controllers

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/repositories"
	"fmt"
)

func GetPosts(ctx *gin.Context) {
	repository := repositories.NewGetPostRepository(DB(ctx))
	posts, _ := repository.Get()
	fmt.Printf("controller resultid: %+v\n", posts)
    ctx.JSON(200, posts)
}
