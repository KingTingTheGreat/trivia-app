package util

import (
	"net/http"
	"trivia-app/api/dlog"
)

const NO_COOKIE = "missing cookie"
const INVALID_TOKEN = "invalid token"
const INVALID_NAME = "repeated name"
const NO_NAME = "no player name provided"
const NOT_FOUND = "player not found"

func Success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func InputError(w http.ResponseWriter, message string) {
	dlog.DLog("writing", message)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
