package auth

import "fmt"

type repository struct {
}

func NewRepository() repository {
	return repository{}
}

func (r repository) CreateUser() error {
	fmt.Println("Call CreateUser in repository")
	return nil
}
