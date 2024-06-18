package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	app.POST("/channels/:id/posts", controllers.CreatePost)
	app.GET("/channel", controllers.GetChannels)
	app.GET("/channels/:id/posts", controllers.GetPosts)

}
