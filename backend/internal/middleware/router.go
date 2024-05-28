package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It w")
	})
	app.GET("/hello", controllers.HelloWorld)
	app.GET("/channels/:channle_id/posts", controllers.GetPosts)
}
