package rest_handlers

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/shared"
	"trivia-app/util"

	"github.com/gorilla/websocket"
)

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	token, err := util.ReadToken(r)
	if err != nil {
		util.InputError(w, r, util.NO_COOKIE)
		return
	}

	player, ok := shared.PlayerStore.GetPlayer(token)
	if !ok {
		util.InputError(w, r, util.INVALID_TOKEN)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		util.InputError(w, r, util.NOT_FOUND)
		return
	}

	if player.Name != name {
		util.InputError(w, r, util.NOT_FOUND)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func PostNewPlayer(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimSpace(r.FormValue("name"))
	dlog.DLog("pName", playerName)
	if util.HasInvalidChar(playerName) {
		dlog.DLog("invalid char", playerName)

		handlers.RenderComponent(w, "error-message.html", util.ErrorData{
			Error: "name can only contain characters a-Z, 0-9, -, and _",
		})
		return
	}

	token, err := util.ReadToken(r)
	if err == nil {
		// recover previous session
		if shared.PlayerStore.VerifyTokenName(token, playerName) {
			dlog.DLog("SESSION RECOVERED")
			util.Redirect(w, r, "/play/"+url.PathEscape(playerName))
			return
		}
	}

	token, err = shared.PlayerStore.InsertPlayer(playerName)
	if err != nil {
		dlog.DLog("repeated name", err.Error())

		handlers.RenderComponent(w, "error-message.html", util.ErrorData{
			Error: err.Error(),
		})
		return
	} else {
		dlog.DLog("NEW PLAYER")
	}

	go func() {
		shared.PlayerListChan <- true
		shared.LeaderboardChan <- true
	}()

	util.WriteToken(w, token)

	dlog.DLog("created new player")
	util.Redirect(w, r, "/play/"+url.PathEscape(playerName))
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	playerName := util.ReadValue(r, "name")
	if playerName == "" {
		util.InputError(w, r, util.NO_NAME)
		return
	}

	token, ok := shared.PlayerStore.NameToToken(playerName)
	if !ok {
		util.InputError(w, r, util.NOT_FOUND)
		return
	}

	action := strings.ToLower(util.ReadValue(r, "action"))
	switch action {
	case "update":
		amountStr := util.ReadValue(r, "amount")
		if amountStr == "" {
			util.InputError(w, r, "no amount provided")
			return
		}
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			util.InputError(w, r, "invalid amount")
			return
		}

		shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{
			ScoreDiff: &amount,
		})
	case "clear":
		shared.PlayerStore.ZeroPlayer(token)
	default:
		dlog.DLog("invalid action", action)
		util.InputError(w, r, util.INVALID_ACTION)
		return
	}

	go func() { shared.LeaderboardChan <- true }()
	go func() {
		player, ok := shared.PlayerStore.GetPlayer(token)
		if !ok {
			return
		}

		var buf bytes.Buffer
		handlers.RenderComponent(&buf, "player-data.html", handlers.Play{Ready: player.ButtonReady, Score: player.Score, ButtonColor: player.ButtonColor})
		if player.Websocket != nil {
			err := player.Websocket.WriteMessage(websocket.TextMessage, buf.Bytes())
			if err != nil {
				player.WsClose <- true
			}
		}
	}()

	util.Success(w, r)
}

func RemovePlayer(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("RemovePlayer()")
	playerName := util.ReadValue(r, "name")
	if playerName == "" {
		dlog.DLog("no name")
		util.InputError(w, r, util.NO_NAME)
		return
	}

	token, ok := shared.PlayerStore.NameToToken(playerName)
	if !ok {
		dlog.DLog("invalid name and token")
		util.InputError(w, r, util.NOT_FOUND)
		return
	}

	player, ok := shared.PlayerStore.GetPlayer(token)
	if !ok {
		dlog.DLog("invalid name and token 2")
		util.InputError(w, r, util.NOT_FOUND)
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

	util.Success(w, r)
}
