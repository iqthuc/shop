package dto

import (
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResult struct {
	UserID           uuid.UUID
	AccessToken      string
	ExpiresIn        time.Duration
	RefreshToken     string
	RefreshExpiresIn time.Duration
	TokenType        string
}

type LoginResponse struct {
	UserID      uuid.UUID     `json:"user_id"`
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	TokenType   string        `json:"token_type"`
}
