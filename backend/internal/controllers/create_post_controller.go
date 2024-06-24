package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createPostRequestParams struct {
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}

func CreatePost(ctx *gin.Context) {
	var post createPostRequestParams
	channelId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.BindJSON(&post); err != nil {
		ctx.String(http.StatusBadRequest, "エラー:invalid json")
		return
	}

	repository := repositories.NewCreatePostRepository(DB(ctx))
	usecase := usecases.NewCreatePostUsecase(repository)
	result, err := usecase.Execute(post.UserId, channelId, post.Content)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.JSON(200, result)
}
