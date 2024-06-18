package infrastructures

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secretKey = []byte("secret")

type CustomClaims struct {
	UserId uint
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "training-web-team-2",
			Subject:   "access_token",
			ID:        uuid.New().String(),
			Audience:  []string{},
		},
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := unsignedToken.SignedString(secretKey)
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := parsedToken.Claims.(*CustomClaims); ok {
		log.Println("hogehogehoge", claims.UserId, claims.ExpiresAt)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
	return nil
}
