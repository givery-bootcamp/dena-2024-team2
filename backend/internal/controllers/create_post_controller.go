package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var post entities.Post
	bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "")
		return
	}

	if err := json.Unmarshal(bytes, &post); err != nil {
		ctx.String(http.StatusBadRequest, "エラー:invalid json")
		return
	}
	fmt.Println("確認", post)

	repository := repositories.NewCreatePostRepository(DB(ctx))
	usecase := usecases.NewCreatePostUsecase(repository)
	result, err := usecase.Execute(post)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.JSON(200, result)
}
