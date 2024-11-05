package handlers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
	"trivia-app/api/util"
)

const COOKIE_NAME = "trivia-app-token"

// only allow alphnum, -, _
var re = regexp.MustCompile("^[a-zA-Z0-9_ -]+$")

func GetPlayerName(w http.ResponseWriter, r *http.Request) {
	token, err := util.ReadToken(r)
	if err != nil {
		util.InputError(w, util.NO_COOKIE)
		return
	}

	player, ok := shared.PlayerStore.GetPlayer(token)
	if !ok {
		util.InputError(w, util.INVALID_TOKEN)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(player.Name))
}

func PostNewPlayer(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimSpace(r.FormValue("name"))
	if playerName == "auth" {
		dlog.DLog("invalid name")
		util.InputError(w, "invalid player name")
		return
	} else if !re.MatchString(playerName) {
		dlog.DLog("invalid char")
		util.InputError(w, "name can only contain characters a-Z, 0-9, -, and _")
		return
	}

	token, err := util.ReadToken(r)
	if err == nil {
		// recover previous session
		if shared.PlayerStore.VerifyTokenName(token, playerName) {
			dlog.DLog("SESSION RECOVERED")
			util.Success(w)
			return
		}
	}

	token, err = shared.PlayerStore.InsertPlayer(playerName)
	if err != nil {
		dlog.DLog("repeated name", err.Error())
		util.InputError(w, err.Error())
		return
	} else {
		dlog.DLog("NEW PLAYER")
	}

	go func() {
		shared.PlayerListChan <- true
		shared.LeaderboardChan <- true
	}()

	util.WriteToken(w, token)
	// dlog.DLog(w)

	dlog.DLog("created new player")
	util.Success(w)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	playerName := r.URL.Query().Get("name")
	if playerName == "" {
		util.InputError(w, util.NO_NAME)
		return
	}

	amountStr := r.URL.Query().Get("amount")
	if amountStr == "" {
		util.InputError(w, "no amount provided")
		return
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		util.InputError(w, "invalid amount")
		return
	}

	token, ok := shared.PlayerStore.NameToToken(playerName)
	if !ok {
		util.InputError(w, util.NOT_FOUND)
		return
	}

	shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{
		ScoreDiff: &amount,
	})

	go func() { shared.LeaderboardChan <- true }()

	util.Success(w)
}

func RemovePlayer(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("RemovePlayer()")
	playerName := r.URL.Query().Get("name")
	if playerName == "" {
		dlog.DLog("no name")
		util.InputError(w, util.NO_NAME)
		return
	}

	token, ok := shared.PlayerStore.NameToToken(playerName)
	if !ok {
		dlog.DLog("invalid name and token")
		util.InputError(w, util.NOT_FOUND)
		return
	}

	player, ok := shared.PlayerStore.GetPlayer(token)
	if !ok {
		dlog.DLog("invalid name and token 2")
		util.InputError(w, util.NOT_FOUND)
		return
	}

	if player.Websocket != nil {
		player.WsClose <- true
	}

	dlog.DLog("deleting player")
	shared.PlayerStore.DeletePlayer(token)

	dlog.DLog("update reader chans")
	go func() {
		shared.PlayerListChan <- true
		shared.BuzzedInChan <- true
		shared.LeaderboardChan <- true
	}()

	util.Success(w)
}
