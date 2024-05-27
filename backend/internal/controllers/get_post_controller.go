package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/repositories"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	repository := repositories.NewGetPostRepository(DB(ctx))
	posts, err := repository.Get()
	fmt.Printf("controller resultid: %+v\n", posts)
	if err != nil {
		handleError(ctx, 500, err)
	} else if posts != nil {
		ctx.JSON(200, posts)
	} else {
		handleError(ctx, 404, errors.New("Not found"))
	}
}
