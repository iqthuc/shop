package auth

import (
	"database/sql"
	"fmt"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateUser(user createUserParams) (createUserResult, error) {
	fmt.Println("Call CreateUser in repository")
	return createUserResult{}, nil
}
