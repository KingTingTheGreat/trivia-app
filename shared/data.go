package shared

import (
	"net/http"
	"slices"
	"strings"
	env "trivia-app"

	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return slices.Contains([]string{
				"localhost:8080",
				"127.0.0.1:8080",
				strings.SplitN(env.GetIP(), "://", 2)[1],
			}, r.Host)
		},
	}
)
