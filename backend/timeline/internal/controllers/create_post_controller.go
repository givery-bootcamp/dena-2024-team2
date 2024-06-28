package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createPostRequestParams struct {
	Content string `json:"content"`
}

func CreatePost(ctx *gin.Context) {
	var post createPostRequestParams
	serverId, err := strconv.Atoi(ctx.Param("serverId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	channelId, err := strconv.Atoi(ctx.Param("channelId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.BindJSON(&post); err != nil {
		ctx.String(http.StatusBadRequest, "エラー:invalid json")
		return
	}

	uid := getUserIdFromContext(ctx)

	createPostRepository := repositories.NewCreatePostRepository(DB(ctx))
	serverRepository := repositories.NewGetServerRepository(DB(ctx))
	channelRepository := repositories.NewGetChannelRepository(DB(ctx))
	usecase := usecases.NewCreatePostUsecase(createPostRepository, serverRepository, channelRepository)
	result, err := usecase.Execute(uid, serverId, channelId, post.Content)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.JSON(200, result)
}
