package page_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
)

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("leaderboard handler")

	data := ws_handlers.LeaderboardData{
		TableId:  "leaderboard",
		Endpoint: "/leaderboard-ws",
	}

	handlers.RenderTemplate(w, "leaderboard.html", data)
}
