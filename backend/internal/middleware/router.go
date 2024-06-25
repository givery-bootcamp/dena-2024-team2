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

	app.POST("/signin", controllers.Signin)
	app.POST("/signup", controllers.Signup)

	authorized := app.Group("/")
	authorized.Use(Auth())
	{
		authorized.GET("/channels", controllers.GetChannels)
		authorized.GET("/channels/:id/posts", controllers.GetPosts)
	}
}
