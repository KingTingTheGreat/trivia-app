package util

import "net/http"

func RequestedHTMX(r *http.Request) bool {
	return r.Header.Get("Hx-Request") != ""
}

func ReadValue(r *http.Request, key string) string {
	value := r.URL.Query().Get(key)
	if value != "" {
		return value
	}
	return r.FormValue(key)
}
