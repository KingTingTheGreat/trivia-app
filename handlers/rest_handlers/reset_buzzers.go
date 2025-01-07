package rest_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/shared"
	"trivia-app/util"
)

func ResetBuzzers(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("reset buzzers")
	// assume authorized bc middleware
	shared.PlayerStore.ResetBuzzers()

	go func() {
		shared.BuzzedInChan <- true
	}()

	dlog.DLog("reset buzzers success")
	util.Success(w, r)
}
