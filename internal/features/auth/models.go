package auth

import (
	"time"

	"github.com/google/uuid"
)

/*
package auth dùng để làm mẫu, còn theo hay không thì tùy @@
[Client]
↓ SignUpRequest --- ↑ signUpResponse
[Handler]
↓ SignUpInput --- ↑ signUpResult
[Usecase]
↓ createUserParams --- ↑ createUserResult
[Repository]
*/

type signUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
type createUserParams struct {
	email        string
	passwordHash string
}

type createUserResult struct {
	email     string
	createdAt time.Time
}

type signUpResult struct {
	email     string
	createdAt time.Time
}

type signUpResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// // login.
type loginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
type loginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int       `json:"expires_in"`
	TokenType    string    `json:"token_type"`
	UserID       uuid.UUID `json:"user_id"`
}
