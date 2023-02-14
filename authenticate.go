package app

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	AuthKey = "_auth_key"
)

type Token struct {
	UserId       string    `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type TokenClaims struct {
	UserId string `json:"uid,omitempty"`
	jwt.RegisteredClaims
}
