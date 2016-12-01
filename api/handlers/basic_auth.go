package handlers

import (
	"net/http"

	"github.com/truetandem/e-QIP-prototype/api/service"
)

func BasicAuth(w http.ResponseWriter, r *http.Request) {

	var respBody struct {
		Username string
		Password string
	}

	if err := DecodeJSON(r.Body, &respBody); err != nil {
		ErrorJSON(w, r, err)
		return
	}

	// Make sure username and password are valid
	account, err := service.Account.BasicAuthentication(respBody.Username, respBody.Password)
	if err != nil {
		ErrorJSON(w, r, err)
		return
	}

	// Generate jwt token
	signedToken, _, err := service.Token.New(account)
	if err != nil {
		ErrorJSON(w, r, err)
		return
	}

	EncodeJSON(w, signedToken)

}
