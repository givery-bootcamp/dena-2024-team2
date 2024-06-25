package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	myConfig "myapp/internal/config"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{myConfig.CorsAllowOrigin}
	config.AllowCredentials = true
	config.AllowMethods = []string{
		"POST",
		"GET",
		"OPTIONS",
	}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Authorization",
	}
	return cors.New(config)
}
