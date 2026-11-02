package page_handlers

import (
	"log"
	"net/http"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
)

type AdminData struct {
	Error       string
	Leaderboard ws_handlers.LeaderboardData
	BuzzedIn    ws_handlers.LeaderboardData
	PlayerList  []string
}

func Admin(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("admin handler")
	playerList := ws_handlers.AllPlayerNames()
	data := AdminData{
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
		PlayerList: playerList,
	}

	log.Println(playerList)

	handlers.RenderTemplate(w, "admin.html", data)
}
