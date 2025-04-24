package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtMaker struct {
	secretKey string
}

func NewJwtMaker(secretKey string) jwtMaker {
	return jwtMaker{
		secretKey: secretKey,
	}
}

func (maker jwtMaker) CreateToken(tkInfo CreateTokenParams) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID:    tkInfo.UserID,
		Role:      tkInfo.Role,
		TokenType: tkInfo.TokenType,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(tkInfo.Duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(maker.secretKey))
}

func (maker jwtMaker) VerifyToken(token string) (*TokenClaims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*TokenClaims)
	if !ok || !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
