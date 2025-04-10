package auth

import (
	"context"
	"fmt"
)

type useCase struct {
	repo Repository
}

func NewUsecase(repo Repository) useCase {
	return useCase{
		repo: repo,
	}
}

type Repository interface {
	CreateUser(user createUserParams) (createUserResult, error)
}

func (u useCase) SignUp(ctx context.Context, input signUpInput) (signUpResult, error) {
	// TODO: hash password, validate uniqueness...
	params := createUserParams{
		Username: input.Username,
		Email:    input.Email, // TODO: hash
		Password: input.Password,
	}
	user, err := u.repo.CreateUser(params)
	if err != nil {
		return signUpResult{}, fmt.Errorf("sign up error: %w", err)
	}
	result := signUpResult{
		Username: user.Username,
		Email:    user.Email,
	}
	return result, nil
}
