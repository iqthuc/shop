package token

import (
	"errors"
	"time"
)

type TokenMaker interface {
	CreateToken(userID, role string, tokenType TokenRole, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenClaims, error)
}

var ErrInvalidToken error = errors.New("token is invalid")
