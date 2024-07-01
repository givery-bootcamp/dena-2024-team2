package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePost(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	uc := usecases.NewDeleteChannelPostUsecase()
	err = uc.Execute(ctx, postId)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			handleError(ctx, 404, errors.New("channel not found"))
		} else {
			handleError(ctx, 500, err)
		}
	} else {
		ctx.JSON(204, "No Content")
	}
}
