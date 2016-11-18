package main

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/linkedin"

	"github.com/gorilla/mux"
)

var (
	oauthStateString     = "random"
	redirectTo           = os.Getenv("API_REDIRECT")
	ErrNoSuchOAuthToken  = errors.New("OAuth Token does not exist")
	ErrInvalidOAuthToken = errors.New("Invalid OAuth Token")
)

func init() {
	// Allow structs to be stored in session
	gob.Register(oauth2.Token{})
}

// Middleware that checks if token is currently in the session and whether
// the token is valid
func OAuthHandler(w http.ResponseWriter, r *http.Request) error {
	log.Println("OAuth2 Middleware")

	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	var token oauth2.Token
	var ok bool

	// Ensure token exists in session
	if token, ok = session.Values["token"].(oauth2.Token); !ok {
		return ErrNoSuchOAuthToken
	}

	// Ensure token is valid (not expired and contains access token)
	if !token.Valid() {
		return ErrInvalidOAuthToken
	}

	log.Println("OAuth2 Token Successfully retrieved from session")

	return nil
}

// authServiceHandler is the initial entry point for authentication.
func authServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]

	config, ok := configureAuthentication(service)
	if !ok {
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
	}

	http.Redirect(w, r, config.AuthCodeURL(oauthStateString), http.StatusTemporaryRedirect)
}

// authCallbackHandler handles responses from the authentication provider.
func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := vars["service"]
	config, ok := configureAuthentication(service)
	if !ok {
		fmt.Printf("Could not determine service with '%s'\n", service)
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
	}

	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("Invalid OAuth state, expected '%s' but recieved '%s'\n", oauthStateString, state)
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, redirectTo, http.StatusTemporaryRedirect)
		return
	}

	session, _ := store.Get(r, sessionName)

	// Github access tokens last forever. Add manual expiry for testing
	token.Expiry = time.Now().Add(time.Minute * 10)

	// Store token in session
	session.Values["token"] = token

	// Persist session data
	session.Save(r, w)

	redirectToWithToken := fmt.Sprintf("%s?token=%s&refresh=%s&expiration=%s", redirectTo, token.AccessToken, token.RefreshToken, token.Expiry)
	http.Redirect(w, r, redirectToWithToken, http.StatusTemporaryRedirect)
}

// configureAuthentication takes a service name and configures the OAuth 2.0 with
// appropriate endpoints and scopes.
func configureAuthentication(service string) (*oauth2.Config, bool) {
	ok := true
	config := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("http://localhost:3000/auth/%s/callback", strings.ToLower(service)),
		ClientID:     os.Getenv(fmt.Sprintf("%s_CLIENT_ID", strings.ToUpper(service))),
		ClientSecret: os.Getenv(fmt.Sprintf("%s_CLIENT_SECRET", strings.ToUpper(service))),
	}

	switch service {
	case "facebook":
		config.Endpoint = facebook.Endpoint
		config.Scopes = []string{
			"public_profile",
			"email",
			"user_birthday",
			"user_work_history",
		}
	case "github":
		config.Endpoint = github.Endpoint
		config.Scopes = []string{
			"user",
		}
	case "google":
		config.Endpoint = google.Endpoint
		config.Scopes = []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		}
	case "linkedin":
		config.Endpoint = linkedin.Endpoint
		config.Scopes = []string{
			"r_basicprofile",
			"r_contactinfo",
			"r_emailaddress",
		}
	// case "microsoft":
	// 	config.Endpoint = microsoft.Endpoint
	// 	config.Scopes = []string{
	// 		"wl.basic",
	// 		"wl.birthday",
	// 		"wl.emails",
	// 	}
	default:
		ok = false
	}

	return config, ok
}
