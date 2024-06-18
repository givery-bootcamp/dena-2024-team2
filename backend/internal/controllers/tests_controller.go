package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/infrastructures"
	"net/http"

	"github.com/gin-gonic/gin"
)

type testRequestParam struct {
	AccessToken string `json:"access_token"`
}

func Test(ctx *gin.Context) {
	var json testRequestParam
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := infrastructures.VerifyToken(json.AccessToken)
	if err != nil {
		if errors.Is(err, entities.ErrNotFound) {
			handleError(ctx, 404, errors.New("channel not found"))
		} else {
			handleError(ctx, 500, err)
		}
	} else {
		ctx.JSON(200, "")
	}
}
