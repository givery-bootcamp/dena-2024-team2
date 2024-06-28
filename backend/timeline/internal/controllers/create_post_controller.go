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

	userId := getUserIdFromContext(ctx)
	userName := getUserNameFromContext(ctx)

	createPostRepository := repositories.NewCreatePostRepository(DB(ctx))
	serverRepository := repositories.NewGetServerRepository(DB(ctx))
	channelRepository := repositories.NewGetChannelRepository(DB(ctx))
	usecase := usecases.NewCreatePostUsecase(createPostRepository, serverRepository, channelRepository)
	result, err := usecase.Execute(serverId, channelId, userId, userName, post.Content)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.JSON(200, newPostConvertToJson(result))
}

func newPostConvertToJson(v *entities.Post) newPostJson {
	return newPostJson{
		Id:        v.Id,
		ChannelId: v.ChannelId,
		UserId:    v.User.Id,
		UserName:  v.User.Name,
		Content:   v.Content,
		CreatedAt: v.CreatedAt,
	}
}

type newPostJson struct {
	Id        int       `json:"id"`
	ChannelId int       `json:"channel_id"`
	UserId    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
