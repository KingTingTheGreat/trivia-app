package handlers

import (
	"net/http"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("reset buzzers")
	// assume authorized bc middleware
	shared.PlayerStore.ResetBuzzers()

	go func() {
		shared.BuzzedInChan <- true
	}()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
