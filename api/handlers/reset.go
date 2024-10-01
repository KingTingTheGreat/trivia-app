package handlers

import (
	"log"
	"net/http"
	"trivia-app/api/shared"
)

func Reset(w http.ResponseWriter, r *http.Request) {
    log.Println("reset buzzers")
    // assume authorized bc middleware 
    shared.PlayerStore.ResetBuzzers()

    go func() {
	shared.BuzzedInChan <- true
    }()

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("success"))
}
