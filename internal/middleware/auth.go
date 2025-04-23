package middleware

import (
	"errors"
	"log/slog"
	"shop/pkg/token"
	"shop/pkg/utils/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeKey    = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
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
			return response.ErrorJson(c, ErrInvalidHeaderFormat, fiber.ErrBadRequest.Code)
		}

		// verify token.
		tokenString := parts[1]
		claims, err := tokenMaker.VerifyToken(tokenString)
		if err != nil || claims.TokenType != token.Access {
			slog.Warn("verify token failed", slog.String("error", err.Error()))
			return response.ErrorJson(c, ErrInvalidToken, fiber.ErrBadRequest.Code)
		}

		c.Locals(AuthorizationPayloadKey, claims)

		return c.Next()
	}
}
