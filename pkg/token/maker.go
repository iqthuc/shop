package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CreateTokenParams struct {
	UserID    uuid.UUID
	Role      string
	TokenType TokenRole
	Duration  time.Duration
}

type TokenMaker interface {
	CreateToken(tkInfo CreateTokenParams) (string, error)
	VerifyToken(token string) (*TokenClaims, error)
}

var ErrInvalidToken error = errors.New("token is invalid")
