package middleware

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/controllers"
)

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It w")
	})
	app.GET("/hello", controllers.HelloWorld)
	app.GET("/posts", controllers.GetPosts)
}
