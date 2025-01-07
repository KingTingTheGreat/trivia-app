package middleware

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// type wrappedWriter struct {
// 	http.ResponseWriter
// 	statusCode int
// }
//
// func (w *wrappedWriter) WriteHeader(statusCode int) {
// 	// w.ResponseWriter.WriteHeader(statusCode)
// 	w.statusCode = statusCode
// }

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// wrapped := &wrappedWriter{
		// 	ResponseWriter: w,
		// 	statusCode:     http.StatusOK,
		// }

		// next.ServeHTTP(wrapped, r)
		// log.Println(wrapped.statusCode, r.Method, r.URL, time.Since(start))
		next.ServeHTTP(w, r)

		log.Println(r.Form)
		body, err := io.ReadAll(r.Body)
		if err == nil {
			log.Printf("Request Body: %s\n", body)
		}

		if !strings.HasPrefix(r.URL.Path, "/public") {
			log.Println(r.Method, r.URL, time.Since(start))
		}
	})
}
