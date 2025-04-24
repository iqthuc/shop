package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenRole int

const (
	Access TokenRole = iota
	Refresh
)

type TokenClaims struct {
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role,omitempty"`
	TokenType TokenRole `json:"type"`
	jwt.StandardClaims
}
