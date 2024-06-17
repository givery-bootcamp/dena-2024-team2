package middleware

import (
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId uint) {
	key := "secret"
	t := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "hoge",
		})
}
