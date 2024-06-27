package usecases

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("db").(*gorm.DB)
}
