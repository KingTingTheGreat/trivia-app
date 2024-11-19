package handlers

import (
	"net/http"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
)

func ResetGame(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("reset buzzers")
	// assume authorized bc middleware
	shared.PlayerStore.ResetGame()

	go func() {
		shared.BuzzedInChan <- true
		shared.LeaderboardChan <- true
		shared.PlayerListChan <- true
	}()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
