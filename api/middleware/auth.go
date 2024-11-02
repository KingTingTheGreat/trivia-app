package middleware

import (
	"net/http"
	"strings"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/auth") {
			next.ServeHTTP(w, r)
			return
		}

		password := r.URL.Query().Get("password")
		if password != shared.Password {
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
