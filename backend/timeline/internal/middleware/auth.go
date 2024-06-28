package middleware

import (
	"myapp/internal/infrastructures"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken, err := ctx.Cookie("training-jwt")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		uid, err := infrastructures.VerifyToken(jwtToken)
		ctx.Set("uid", uid)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Next()
	}
}
