package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func DB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("db").(*gorm.DB)
}

//nolint:unused
func handleError(ctx *gin.Context, status int, err error) {
	res := ErrorResponse{
		Message: err.Error(),
	}
	ctx.JSON(status, &res)
}

func getUserIdFromContext(ctx *gin.Context) int {
	return ctx.MustGet("uid").(int)
}
