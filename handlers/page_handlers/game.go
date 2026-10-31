package page_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
)

type GameData struct {
	Leaderboard ws_handlers.LeaderboardData
	BuzzedIn    ws_handlers.LeaderboardData
}

func Game(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("game handler")

	data := GameData{
		Leaderboard: ws_handlers.LeaderboardData{
			TableId:     "leaderboard",
			Title:       "Leaderboard",
			Description: "Top players by score",
			Endpoint:    "/leaderboard-ws",
		},
		BuzzedIn: ws_handlers.LeaderboardData{
			TableId:     "buzzed-in",
			Title:       "Buzzed In",
			Description: "Players ranked by reaction speed",
			Endpoint:    "/buzzed-in-ws",
		},
	}

	handlers.RenderTemplate(w, "game.html", data)
}
