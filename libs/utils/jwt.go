package utils

import (
	"app"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"os"
	"time"
)

// https://pkg.go.dev/github.com/golang-jwt/jwt/v4#readme-examples

func CreateJWT(userId string, timeout time.Duration) (string, error) {

	expires := time.Now().Add(timeout)

	claims := app.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   "jwt",
			Audience:  []string{"client"},
			ExpiresAt: jwt.NewNumericDate(expires),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// access token
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func ParseJWT(accessToken string) (*app.TokenClaims, error) {

	token, err := jwt.ParseWithClaims(accessToken, &app.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*app.TokenClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("error accessToken")
	}
}
