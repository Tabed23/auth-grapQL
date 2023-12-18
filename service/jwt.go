package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtSecert = []byte("my_secret_key")

// Generate a JWT token from the username and email
func GenrarteJwtToken(ctx context.Context, userID, email string) (string, error) {
	j := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaims{
		ID:    userID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := j.SignedString(jwtSecert)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Validate the token
func ValidateJwt(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecert, nil
	})
}
