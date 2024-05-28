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
	channelId, err := strconv.Atoi(ctx.Param("channel_id"))
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
			Title:     v.Title,
			Body:      v.Body,
			UserName:  v.User.Name,
			UserId:    v.User.Id,
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
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserId    int       `json:"userId"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
