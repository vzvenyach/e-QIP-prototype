package main

import (
	"fmt"
	"log"
	"net/http"
)

func LoggerHandler(w http.ResponseWriter, r *http.Request) error {
	log.Println("Log handler middleware")

	// Add eqip api information to all requests in headers
	w.Header().Set("X-Eqip-Media-Type", fmt.Sprintf("%s.%s", APIName, APIVersion))

	return nil
}
