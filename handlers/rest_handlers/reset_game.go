package rest_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/shared"
	"trivia-app/util"
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

	util.Success(w, r)
}
