package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"myproject/web/model"
)

var jwtSecret = []byte("xxoxx")


type Claims struct {
	UserID uint `json:"user_id"`
	Username string	`json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.User) (string, error) {

	claims := Claims{
		UserID: user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "go-gin-app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token)(interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}