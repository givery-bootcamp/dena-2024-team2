package infrastructures

import (
	"fmt"
	"log"
	"myapp/internal/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	claims := CustomClaims{
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
	signedToken, err := unsignedToken.SignedString([]byte(config.JWTSecret))
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(token string) (int, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWTSecret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := parsedToken.Claims.(*CustomClaims); ok {
		uid, err := strconv.Atoi(claims.Subject)
		if err != nil {
			return 0, err
		}
		return uid, nil
	} else {
		log.Fatal("unknown claims type, cannot proceed")
		return 0, fmt.Errorf("failed to convert token to customClaims")
	}
}
