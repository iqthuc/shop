package token

import (
	"github.com/golang-jwt/jwt"
)

type TokenRole int

const (
	Access TokenRole = iota
	Refresh
)

type TokenClaims struct {
	UserID    string    `json:"user_id"`
	Role      string    `json:"role,omitempty"`
	TokenType TokenRole `json:"type"`
	jwt.StandardClaims
}
