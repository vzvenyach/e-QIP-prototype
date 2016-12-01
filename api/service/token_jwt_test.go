package service

import (
	"testing"

	"github.com/truetandem/e-QIP-prototype/api/model"
)

func TestNewToken(t *testing.T) {
	a := model.Account{
		ID: 1,
	}

	tokenService := NewJwtTokenService()
	token, _, err := tokenService.New(a)
	if err != nil {
		t.Fatal(err)
	}

	valid, _ := tokenService.Valid(token)
	if !valid {
		t.Fatalf("Expected Jwt Token to be valid")
	}

	valid, _ = tokenService.Valid("badtoken")
	if valid {
		t.Fatalf("Expected Jwt Token to be invalid")
	}
}
