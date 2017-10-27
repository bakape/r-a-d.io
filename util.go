package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// Text-only 404 response
func text404(w http.ResponseWriter) {
	http.Error(w, "404 not found", 404)
}

// Text-only 400 response
func text400(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("400 %s", err), 400)
}

func text403(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("403 %s", err), 403)
}

// Text-only 500 response
func text500(w http.ResponseWriter, r *http.Request, err interface{}) {
	http.Error(w, fmt.Sprintf("500 %s", err), 500)
	logError(r, err)
}

// Log an error together with the client's IP and stack trace
func logError(r *http.Request, err interface{}) {
	log.Printf("server: %s\n%s\n", err, debug.Stack())
}
