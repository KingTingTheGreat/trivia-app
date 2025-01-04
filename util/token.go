package util

import "net/http"

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
		Name:  COOKIE_NAME,
		Value: token,
	}

	http.SetCookie(w, &cookie)
}
