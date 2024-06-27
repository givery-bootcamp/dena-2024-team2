package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type updatePostRequestParams struct {
	Content string `json:"content"`
}

func UpdatePost(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	var req updatePostRequestParams
	if err := ctx.BindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	postRepository := repositories.NewPostRepository(DB(ctx))
	uc := usecases.NewUpdatePostUsecase(postRepository)
	uid := getUserIdFromContext(ctx)
	post, err := uc.Execute(ctx, postId, uid, req.Content)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, post)
}
