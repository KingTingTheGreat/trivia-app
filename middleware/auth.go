package middleware

import (
	"net/http"
	"strings"
	"trivia-app/dlog"
	"trivia-app/shared"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/auth") {
			next.ServeHTTP(w, r)
			return
		}

		password := r.URL.Query().Get("password")
		if strings.TrimSpace(password) == "" {
			dlog.DLog("unauthorized. no password")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("password is required"))
			return
		} else if password != shared.Password {
			dlog.DLog("unauthorized", "input", password, "stored", shared.Password)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("incorrect password"))
			return
		}

		dlog.DLog("auth success")
		next.ServeHTTP(w, r)
		return
	})
}
