package rest_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/shared"
)

func ResetBuzzers(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("reset buzzers")
	// assume authorized bc middleware
	shared.PlayerStore.ResetBuzzers()

	go func() {
		shared.BuzzedInChan <- true
	}()

	dlog.DLog("reset buzzers success")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
