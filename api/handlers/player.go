package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"trivia-app/api/shared"
	"trivia-app/api/util"
)

const COOKIE_NAME = "trivia-app-token"

func GetPlayerName(w http.ResponseWriter, r *http.Request) {
    token, err := util.ReadToken(r)
    if err != nil {
	util.UserInputError(w, "no cookie")
	return
    }

    player, ok := shared.PlayerStore.GetPlayer(token)
    if !ok {
	util.UserInputError(w, "invalid token")
	return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(player.Name))
}

func PostNewPlayer(w http.ResponseWriter, r *http.Request) {
    playerName := strings.TrimSpace(r.FormValue("name"))
    if playerName == "auth" {
	log.Println("invalid name")
	util.UserInputError(w, "invalid player name")
	return
    }

    token, err := util.ReadToken(r)
    if err == nil {
	// recover previous session
	if shared.PlayerStore.VerifyTokenName(token, playerName) {
	    log.Println("session recovered")
	    w.WriteHeader(http.StatusOK)
	    w.Write([]byte("success"))
	    return
	}
    }

    token, err = shared.PlayerStore.InsertPlayer(playerName)
    if err != nil {
	log.Println("repeated name")
	util.UserInputError(w, err.Error())
	return
    }

    go func() {
	shared.PlayerListChan <- true
	shared.LeaderboardChan <- true
	shared.BuzzedInChan <- true 
    }()

    util.WriteToken(w, token)

    log.Println("created new player")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("success"))
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
    playerName := r.URL.Query().Get("name")
    if playerName == "" {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("no name provided"))
	return
    }

    amountStr := r.URL.Query().Get("amount")
    if amountStr == "" {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("no amount provided"))
	return
    }
    amount, err := strconv.Atoi(amountStr)
    if err != nil {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("invalid amount"))
	return
    }


    token, ok := shared.PlayerStore.NameToToken(playerName)
    if !ok {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("player not found"))
	return
    }

    shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{
	ScoreDiff: &amount,
    })
}

func RemovePlayer(w http.ResponseWriter, r *http.Request) {
    playerName := r.URL.Query().Get("name")
    if playerName == "" {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("no name provided"))
	return
    }

    token, ok := shared.PlayerStore.NameToToken(playerName)
    if !ok {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("player not found"))
	return
    }

    shared.PlayerStore.DeletePlayer(token)

    go func() {
	shared.PlayerListChan <- true 
	shared.BuzzedInChan <- true
	shared.LeaderboardChan <- true
    }()

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("success"))
}
