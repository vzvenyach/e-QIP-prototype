package service

import (
	"log"

	"github.com/truetandem/e-QIP-prototype/api/db"
	"github.com/truetandem/e-QIP-prototype/api/model"
	pg "gopkg.in/pg.v5"
)

// AccountPgService provides Account actions using a Postgres database
type AccountPgService struct {
	db *pg.DB
}

// BasicAuth checks if the username and password are valid and returns the users account
func (a *AccountPgService) BasicAuthentication(username, password string) (account model.Account, err error) {
	var basicMembership model.BasicAuthMembership

	// Find if basic auth record exists for given account username
	err = a.db.Model(&basicMembership).
		Column("basic_auth_membership.*", "Account").
		Where("Account.username = ?", username).
		Select()

	if err != nil {
		log.Println(err)
		return account, ErrAccoundDoesNotExist
	}

	// Check if plaintext password matches hashed password
	if matches := basicMembership.PasswordMatch(password); !matches {
		return account, ErrPasswordDoesNotMatch
	}

	return *basicMembership.Account, nil
}

func NewAccountPgService() *AccountPgService {
	db := db.NewDB()
	return &AccountPgService{db}
}
