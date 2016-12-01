package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/truetandem/e-QIP-prototype/api/db"
	"github.com/truetandem/e-QIP-prototype/api/model"
)

func TestAccountPgBasicAuthentication(t *testing.T) {
	db := db.NewDB()
	username := fmt.Sprintf("user-%v", time.Now().Unix())
	pw := "admin"

	a := model.Account{
		Username:  username,
		Firstname: "Admin",
		Lastname:  "Last",
	}

	err := db.Insert(&a)
	if err != nil {
		t.Fatal(err)
	}

	basic := model.BasicAuthMembership{
		AccountID: a.ID,
	}
	basic.HashPassword(pw)

	err = db.Insert(&basic)
	if err != nil {
		t.Fatal(err)
	}

	accountService := AccountPgService{db}
	_, err = accountService.BasicAuthentication(a.Username, pw)
	if err != nil {
		t.Fatalf("Expected password to match\n")
	}

	_, err = accountService.BasicAuthentication(a.Username, "incorrect-password")
	if err == nil {
		t.Fatalf("Expected incorrect password\n")
	}
}
