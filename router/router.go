package router

import (
	"net/http"
	"trivia-app/handlers/page_handlers"
	"trivia-app/handlers/rest_handlers"
	"trivia-app/handlers/ws_handlers"
	"trivia-app/shared"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	// serve files (css, js, mp3, etc)
	router.Handle("/public/", http.StripPrefix("/public/",
		http.FileServer(http.Dir("./public"))),
	)

	// pages
	router.HandleFunc("/", page_handlers.Home)
	router.HandleFunc("GET /play/{name}", page_handlers.Play)
	router.HandleFunc("GET /join", page_handlers.Join)
	router.HandleFunc("GET /leaderboard", page_handlers.Leaderboard)
	router.HandleFunc("GET /buzzed-in", page_handlers.BuzzedIn)
	router.HandleFunc("GET /game", page_handlers.Game)
	router.HandleFunc("GET /pages", page_handlers.Pages)

	router.HandleFunc("GET /health", rest_handlers.Health)

	// player flow, create and buzz
	router.HandleFunc("POST /player", rest_handlers.PostNewPlayer)
	if shared.ReactiveBuzzers() {
		router.HandleFunc("/buzz", ws_handlers.BuzzWS)
	} else {
		router.HandleFunc("POST /buzz", rest_handlers.Buzz)
	}

	// game controls
	router.HandleFunc("GET /control", page_handlers.Control)
	router.HandleFunc("POST /auth/reset-buzzers", rest_handlers.ResetBuzzers)
	router.HandleFunc("POST /auth/reset-game", rest_handlers.ResetGame)
	router.HandleFunc("PUT /auth/player", rest_handlers.UpdatePlayer)
	router.HandleFunc("DELETE /auth/player", rest_handlers.RemovePlayer)

	// data retrieval websockets
	router.HandleFunc("/leaderboard-ws", ws_handlers.LeaderboardWS)
	router.HandleFunc("/buzzed-in-ws", ws_handlers.BuzzedInWS)
	router.HandleFunc("/players-ws", ws_handlers.PlayerListWS)

	return router
}
