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
		authorized.GET("/servers/:serverId/channels", controllers.GetChannels)
		authorized.POST("/servers/:serverId/channels", controllers.CreateChannels)
		authorized.GET("/channels/:channelId/posts", controllers.GetPosts)
		authorized.POST("/servers/:serverId/channels/:channelId/posts", controllers.CreatePost)
		authorized.POST("/servers", controllers.CreateServer)
	}
}
