package token

import "time"

type TokenMaker interface {
	CreateAccessToken(userID, role string, tokenType TokenType, duration time.Duration) (string, error)
	VerifyAccessToken(token string) (*TokenClaims, error)
}
