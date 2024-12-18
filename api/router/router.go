package router

import (
	"net/http"
	"trivia-app/api/handlers"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	router.HandleFunc("GET /api/player", handlers.GetPlayerName)
	router.HandleFunc("POST /api/player", handlers.PostNewPlayer)

	router.HandleFunc("GET /api/verify", handlers.Verify)

	router.HandleFunc("/api/buzz", handlers.BuzzWs)

	router.HandleFunc("POST /api/auth/reset", handlers.Reset)

	router.HandleFunc("POST /api/auth/reset-game", handlers.ResetGame)

	router.HandleFunc("PUT /api/auth/player", handlers.UpdatePlayer)
	router.HandleFunc("PUT /api/auth/clear-player", handlers.ClearPlayer)
	router.HandleFunc("DELETE /api/auth/player", handlers.RemovePlayer)

	router.HandleFunc("/api/leaderboard", handlers.Leaderboard)
	router.HandleFunc("/api/buzzed-in", handlers.BuzzedIn)
	router.HandleFunc("/api/players", handlers.PlayerList)

	return router
}
