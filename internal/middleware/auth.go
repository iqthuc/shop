package middleware

import (
	"errors"
	"shop/pkg/token"
	"shop/pkg/utils/response"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeKey    = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

var (
	ErrHeaderNotProvided   = errors.New("authorization header is not provided")
	ErrInvalidHeaderFormat = errors.New("authorization header format is invalid")
	ErrInvalidToken        = errors.New("token is invalid")
)

func JWTAuth(tokenMaker token.TokenMaker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(authorizationHeaderKey)
		if authHeader == "" {
			return response.ErrorJson(c, ErrHeaderNotProvided, fiber.ErrBadRequest.Code)
		}

		// check token type.
		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != authorizationTypeKey {
			log.Error(parts)
			return response.ErrorJson(c, ErrInvalidHeaderFormat, fiber.ErrBadRequest.Code)
		}

		// verify token.
		tokenString := parts[1]
		claims, err := tokenMaker.VerifyToken(tokenString)
		if err != nil || claims.TokenType != token.Access {
			return response.ErrorJson(c, ErrInvalidToken, fiber.ErrBadRequest.Code)
		}
		c.Locals(authorizationPayloadKey, claims)

		return c.Next()
	}
}
