package token

import (
	errs "shop/pkg/utils/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) TokenMaker {
	return JWTMaker{
		secretKey: secretKey,
	}
}
func (maker JWTMaker) CreateToken(userID, role string, tokenRole TokenRole, duration time.Duration) (string, error) {
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

func (maker JWTMaker) VerifyToken(token string) (*TokenClaims, error) {
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
