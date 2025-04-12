package auth

import "time"

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
