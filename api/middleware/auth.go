package middleware

import (
	"log"
	"net/http"
	"strings"
	"trivia-app/api/shared"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api/auth") {
	    log.Println("not an auth route")
	    next.ServeHTTP(w, r)
	    return
	}

	password := r.URL.Query().Get("password")
	if password != shared.Password {
    log.Println("unauthorized", "input", password, "stored", shared.Password)
	    w.WriteHeader(http.StatusUnauthorized)
	    w.Write([]byte("incorrect password"))
	    return
	}

	log.Println("auth success")
	next.ServeHTTP(w, r)
	return
    })
}
