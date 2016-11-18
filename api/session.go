package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store       = sessions.NewCookieStore([]byte("muy-secret"))
	sessionName = "eqip"
)

func SessionHandler(w http.ResponseWriter, r *http.Request) error {
	log.Println("Session handler middleware")

	// Create or obtain existing session
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	if session.IsNew {
		// Do/add stuff when new session is created
	}

	return nil
}
