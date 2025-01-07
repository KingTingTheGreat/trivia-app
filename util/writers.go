package util

import (
	"fmt"
	"net/http"
	"trivia-app/dlog"
	"trivia-app/handlers"
)

const NO_COOKIE = "missing cookie"
const INVALID_TOKEN = "invalid token"
const INVALID_NAME = "repeated name"
const INVALID_ACTION = "invalid action"
const NO_NAME = "no player selected"
const NOT_FOUND = "player not found"

func Success(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if !RequestedHTMX(r) {
		w.Write([]byte("success"))
	}
}

func RedirectError(w http.ResponseWriter, r *http.Request, message string) {
	dlog.DLog("redirecting error", message)

	var redirectUrl string
	if message != "" {
		redirectUrl = fmt.Sprintf("/?error=%s", message)
	} else {
		redirectUrl = "/"
	}

	w.Header().Set("HX-Location", redirectUrl)
	w.Header().Set("Location", redirectUrl)

	if RequestedHTMX(r) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
	}
}

type ErrorData struct {
	Error string
}

func InputError(w http.ResponseWriter, r *http.Request, message string) {
	dlog.DLog("redirecting error", message)

	if RequestedHTMX(r) {
		handlers.RenderComponent(w, "error-message.html", ErrorData{Error: message})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(message))
	}
}
