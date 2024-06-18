package infrastructures

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UserId uint
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	key := []byte("secret")
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
	signedToken, err := unsignedToken.SignedString(key)
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}

	return signedToken, nil
}
