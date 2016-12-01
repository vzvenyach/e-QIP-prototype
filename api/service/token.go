package service

import (
	"time"

	"github.com/truetandem/e-QIP-prototype/api/model"
)

// TokenService is an interface for a service that can generate new tokens and validate them
type TokenService interface {
	New(model.Account) (string, time.Time, error)
	Valid(string) (bool, error)
}
