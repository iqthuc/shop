package password

import (
	"testing"
)

func TestHashAndCheckPassword_Success(t *testing.T) {
	password := "myStrongPassword123!"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error while hashing, got: %v", err)
	}

	if len(hashed) == 0 {
		t.Fatal("expected hashed password, got empty string")
	}

	if !CheckPasswordHash(password, hashed) {
		t.Error("expected password to match hashed password, but it didn't")
	}
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	password := "correctPassword"
	wrongPassword := "wrongPassword"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	if CheckPasswordHash(wrongPassword, hashed) {
		t.Error("expected wrong password to not match hash, but it did")
	}
}
