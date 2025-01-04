package middleware

import "net/http"

func CreateStack(xs ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func Stack() func(http.Handler) http.Handler {
	return CreateStack(
		Logger,
		Auth,
		UserInfo,
	)
}
