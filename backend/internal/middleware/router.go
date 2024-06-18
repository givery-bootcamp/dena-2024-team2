package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.POST("/test", controllers.Test)
	app.POST("/signin", controllers.Signin)

	authorized := app.Group("/")
	authorized.Use(Auth)
	{
		authorized.GET("/channel", controllers.GetChannels)
		authorized.GET("/channels/:id/posts", controllers.GetPosts)
	}
}
