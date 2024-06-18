package middleware

import (
	"fmt"
	"log"
	"myapp/internal/infrastructures"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken, err := extractBearerToken(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		log.Printf("ninsyoutyantodekiteiru")
		err = infrastructures.VerifyToken(jwtToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err)
			return
		}
		ctx.Next()
	}
}

func extractBearerToken(headerString string) (string, error) {
	if headerString == "" {
		return "", fmt.Errorf("required header is empty")
	}

	jwtToken := strings.Split(headerString, " ")
	if len(jwtToken) != 2 {
		return "", fmt.Errorf("authorization header requires Bearer before token: see https://jwt.io/introduction")
	}

	return jwtToken[1], nil
}
