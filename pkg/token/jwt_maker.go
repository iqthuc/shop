package token

import (
	errs "shop/pkg/utils/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtMaker struct {
	secretKey string
}

//nolint:ireturn
func NewJwtMaker(secretKey string) jwtMaker {
	return jwtMaker{
		secretKey: secretKey,
	}
}

func (maker jwtMaker) CreateToken(userID, role string, tokenRole TokenRole, duration time.Duration) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID:    userID,
		Role:      role,
		TokenType: tokenRole,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(duration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(maker.secretKey))
}

func (maker jwtMaker) VerifyToken(token string) (*TokenClaims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errs.ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, keyFunc)
	if err != nil {
		return nil, errs.ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*TokenClaims)
	if !ok || !jwtToken.Valid {
		return nil, errs.ErrInvalidToken
	}

	return claims, nil
}
