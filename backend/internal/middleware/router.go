package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})

	app.POST("/signin", controllers.Signin)
	app.POST("/signup", controllers.Signup)
	authorized := app.Group("/")
	authorized.Use(Auth())
	{
		authorized.GET("/servers", controllers.GetServers)
		authorized.GET("/servers/:id/channels", controllers.GetChannels)
		authorized.POST("/servers/:id/channels", controllers.CreateChannels)
		authorized.GET("/channels/:id/posts", controllers.GetPosts)
		authorized.POST("/channels/:id/posts", controllers.CreatePost)
		authorized.POST("/servers", controllers.CreateServer)
	}
}
