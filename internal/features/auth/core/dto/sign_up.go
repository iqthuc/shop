package dto

import "time"

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignUpInput struct {
	Email    string
	Password string
}

type SignUpResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
