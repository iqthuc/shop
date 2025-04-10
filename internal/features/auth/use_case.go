package auth

import (
	"context"
	"fmt"
)

type Repository interface {
	CreateUser() error
}

type useCase struct {
	repo Repository
}

func NewUsecase(repo Repository) useCase {
	return useCase{
		repo: repo,
	}
}

func (u useCase) Login(ctx context.Context, email string, password string) error {
	fmt.Println("call Login in use case")
	u.repo.CreateUser()
	return nil
}
