package ws_handlers

import (
	"bytes"
	"log"
	"net/http"
	"time"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/shared"

	"github.com/gorilla/websocket"
)

var playerListWS = shared.NewWebsocketStore()

type PlayerListData struct {
	PlayerList []string
}

func PlayerListWS(w http.ResponseWriter, r *http.Request) {
	log.Println("playerlistws()")
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading player list connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	playerListWS.InsertConn(conn)

	// write current state (this is for reconnections)
	var buf bytes.Buffer
	handlers.RenderComponent(&buf, "playerlist-dropdown.html", PlayerListData{
		PlayerList: AllPlayerNames(),
	})
	conn.WriteMessage(websocket.TextMessage, buf.Bytes())

	go playerListWS.KeepAlive(conn)
}

func AllPlayerNames() []string {
	allPlayers := shared.PlayerStore.AllPlayers()

	allPlayerNames := make([]string, len(allPlayers))
	for i, player := range allPlayers {
		allPlayerNames[i] = player.Name
	}

	return allPlayerNames
}

func BroadcastPlayerList() {
	var buf bytes.Buffer
	for range shared.PlayerListChan {
		now := time.Now()
		for time.Since(now) < 50*time.Millisecond {
			select {
			case <-shared.PlayerListChan:
			default:
				break
			}
		}
		dlog.DLog("player list chan", AllPlayerNames())
		handlers.RenderComponent(&buf, "playerlist-dropdown.html", PlayerListData{
			PlayerList: AllPlayerNames(),
		})
		playerListWS.WriteToAll(buf.Bytes())
	}
}
