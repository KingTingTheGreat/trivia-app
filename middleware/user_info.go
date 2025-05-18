package middleware

import (
	"context"
	"net/http"
	"trivia-app/shared"
	"trivia-app/util"
)

var paths []string = []string{
	"/player",
}

func UserInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := util.ReadToken(r)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		player, ok := shared.PlayerStore.GetPlayer(token)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)
		ctx = context.WithValue(ctx, "player", &player)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
