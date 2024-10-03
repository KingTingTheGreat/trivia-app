package handlers

import (
	"log"
	"net/http"
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
		log.Println("error upgrading player list connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a := shared.PlayerStore.AllNamesTokens()
	log.Println("sending", a)
	conn.WriteJSON(shared.PlayerStore.AllNamesTokens())
	playerListWS.InsertConn(conn)

	go playerListWS.KeepAlive(conn)
}

func BroadcastPlayerList() {
	for range shared.PlayerListChan {
		log.Println("player list chan")
		playerList := shared.PlayerStore.AllNamesTokens()
		playerListWS.WriteToAll(playerList)
	}
}
