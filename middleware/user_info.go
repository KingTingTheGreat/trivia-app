package middleware

import (
	"context"
	"log"
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
			log.Println("no cookie")
			next.ServeHTTP(w, r)
			return
		}

		player, ok := shared.PlayerStore.GetPlayer(token)
		if !ok {
			log.Println("invalid token")
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)
		ctx = context.WithValue(ctx, "player", &player)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
