package controllers

import (
	"errors"
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPosts(ctx *gin.Context) {
	repository := repositories.NewGetPostRepository(DB(ctx))
	posts, err := repository.Get()
	fmt.Printf("controller resultid: %+v\n", posts)
	postsJson := convertToJson(posts)
	if err != nil {
		handleError(ctx, 500, err)
	} else if posts != nil {
		ctx.JSON(200, postsJson)
	} else {
		handleError(ctx, 404, errors.New("Not found"))
	}
}

func convertToJson(posts []entities.Posts) ChannnelPostJson {
	entityPosts := make([]ChannnelPostJson, len(posts))
	for i, v := range posts {
		entityPosts[i] = ChannnelPostJson{
			Id:        v.Id,
			Title:     v.Title,
			Body:      v.Body,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			DeletedAt: v.DeletedAt,
		}
	}
	return entityPosts
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
