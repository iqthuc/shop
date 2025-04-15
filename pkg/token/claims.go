package token

import (
	"github.com/golang-jwt/jwt"
)

type TokenType int

const (
	Access TokenType = iota
	Refresh
)

type TokenClaims struct {
	UserID    string    `json:"user_id"`
	Role      string    `json:"role,omitempty"`
	TokenType TokenType `json:"type"`
	jwt.StandardClaims
}
