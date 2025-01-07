package middleware

import (
	"net/http"
	"strings"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/shared"
	"trivia-app/util"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/auth") {
			next.ServeHTTP(w, r)
			return
		}

		r.ParseForm()
		password := strings.TrimSpace(util.ReadValue(r, "password"))
		if password == "" {
			dlog.DLog("unauthorized. no password")
			if util.RequestedHTMX(r) {
				handlers.RenderComponent(w, "error-message.html", util.ErrorData{
					Error: "password is required",
				})
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("password is required"))
			}
			return
		} else if password != shared.Password {
			dlog.DLog("unauthorized", "input", password, "stored", shared.Password)
			if util.RequestedHTMX(r) {
				handlers.RenderComponent(w, "error-message.html", util.ErrorData{
					Error: "incorrect password",
				})
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("incorrect password"))
			}
			return
		}

		dlog.DLog("auth success")
		next.ServeHTTP(w, r)
		return
	})
}
