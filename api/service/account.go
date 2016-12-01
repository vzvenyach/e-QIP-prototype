package service

import (
	"errors"

	"github.com/truetandem/e-QIP-prototype/api/model"
)

var (
	ErrPasswordDoesNotMatch = errors.New("Password does not match")
	ErrAccoundDoesNotExist  = errors.New("Account does not exist")
)

// AccountService is a service that performs account related actions
type AccountService interface {
	BasicAuthentication(username, password string) (model.Account, error)
}
