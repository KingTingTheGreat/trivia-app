package util

import (
	"net/http"
	"time"
)

const COOKIE_NAME = "trivia-app-token"

func ReadToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func WriteToken(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:    COOKIE_NAME,
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour * 365), // one year expiry
	}

	http.SetCookie(w, &cookie)
}
