package handlers

import (
	"net/http"
	"time"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
)

var playerListWS = shared.NewWebsocketStore()

type playerListPlayer struct {
	Name  string
	Token string
}

func PlayerList(w http.ResponseWriter, r *http.Request) {
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading player list connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a := shared.PlayerStore.AllNamesTokens()
	dlog.DLog("sending", a)
	conn.WriteJSON(shared.PlayerStore.AllNamesTokens())
	playerListWS.InsertConn(conn)

	go playerListWS.KeepAlive(conn)
}

func BroadcastPlayerList() {
	for range shared.PlayerListChan {
		now := time.Now()
		for time.Since(now) < 50*time.Millisecond {
			select {
			case <-shared.PlayerListChan:
			default:
				break
			}
		}
		dlog.DLog("player list chan")
		playerList := shared.PlayerStore.AllNamesTokens()
		playerListWS.WriteToAll(playerList)
	}
}
