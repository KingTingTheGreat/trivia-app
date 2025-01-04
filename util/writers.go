package util

import (
	"fmt"
	"net/http"
	"trivia-app/dlog"
)

const NO_COOKIE = "missing cookie"
const INVALID_TOKEN = "invalid token"
const INVALID_NAME = "repeated name"
const INVALID_ACTION = "invalid action"
const NO_NAME = "no player selected"
const NOT_FOUND = "player not found"

func Success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func RedirectError(w http.ResponseWriter, r *http.Request, message string) {
	dlog.DLog("redirecting error", message)
	if message != "" {
		http.Redirect(w, r, fmt.Sprintf("/?error=%s", message), http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func InputError(w http.ResponseWriter, message string) {
	dlog.DLog("redirecting error", message)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
