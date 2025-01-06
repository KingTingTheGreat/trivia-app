package page_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
)

func BuzzedIn(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("buzzed in handler")

	data := ws_handlers.LeaderboardData{
		Title:    "Buzzed In",
		TableId:  "buzzed-in",
		Endpoint: "/buzzed-in-ws",
	}

	handlers.RenderTemplate(w, "buzzed-in.html", data)
}
