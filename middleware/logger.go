package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		if !strings.HasPrefix(r.URL.Path, "/public") {
			log.Println(r.Method, r.URL, time.Since(start))
		} else {
			log.Println(r.Method, r.URL, time.Since(start))
		}
	})
}
