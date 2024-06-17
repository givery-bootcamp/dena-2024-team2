package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	channelId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	uc := usecases.NewGetChannelPostUsecase()
	posts, err := uc.Execute(ctx, channelId)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			handleError(ctx, 404, errors.New("channel not found"))
		} else {
			handleError(ctx, 500, err)
		}
	} else {
		ctx.JSON(200, convertToJson(posts))
	}
}

func convertToJson(posts entities.Posts) ResponseJson {
	jsonPosts := make([]ChannnelPostJson, len(posts))
	for i, v := range posts {
		jsonPosts[i] = ChannnelPostJson{
			Id:        v.Id,
			ChannelId: v.ChannelId,
			UserId:    v.UserId,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
	}
	return ResponseJson{Posts: jsonPosts}
}

type ResponseJson struct {
	Posts []ChannnelPostJson `json:"posts"`
}

type ChannnelPostJson struct {
	Id        int        `json:"id"`
	ChannelId int        `json:"channel_id"`
	UserId    int        `json:"user_id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
