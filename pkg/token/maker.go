package token

import "time"

type TokenMaker interface {
	CreateToken(userID, role string, tokenType TokenRole, duration time.Duration) (string, error)
	VerifyToken(token string) (*TokenClaims, error)
}
