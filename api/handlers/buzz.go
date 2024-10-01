package handlers

import (
	"log"
	"net/http"
	"trivia-app/api/shared"
	"trivia-app/api/util"

	"github.com/gorilla/websocket"
)

func BuzzWs(w http.ResponseWriter, r *http.Request) {
	log.Println("web socket")
	// read cookie
	token, err := util.ReadToken(r)
	if err != nil {
		util.UserInputError(w, "no cookie")
		return
	}

	// player associated with cookie
	player, ok := shared.PlayerStore.GetPlayer(token)
	if !ok {
		log.Println("player does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// upgrade to websocket
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading connection to websocket")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	go websocketHandler(conn, token, player)
}

func websocketHandler(conn *websocket.Conn, token string, player shared.Player) {
	defer func(conn *websocket.Conn) {
		log.Println("closing websocket", player.Name)
		shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{Websocket: nil})
		err := conn.Close()
		if err != nil {
			log.Println("error closing websocket connection")
		}
	}(conn)

	// first message is used to verify player, not buzz in
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println("could not write to websocket")
		return
	}
	name := string(p)
	if player.Name != name {
		log.Println("name and token do not match")
		return
	}

	// update player store entry with websocket
	shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{Websocket: conn})

	// update client-side with correct button state
	if player.ButtonReady {
		log.Println(player.Name, "ready")
		conn.WriteMessage(websocket.TextMessage, []byte("ready"))
	} else {
		log.Println(player.Name, "buzz")
		conn.WriteMessage(websocket.TextMessage, []byte("buzz"))
	}

	go func() { shared.LeaderboardChan <- true }()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("websocket error", player.Name)
			break
		}
		name := string(p)
		log.Println("websocket message:", name)
		if shared.PlayerStore.BuzzIn(token, name) {
			log.Println(name, "buzzed in")
			conn.WriteMessage(websocket.TextMessage, []byte("buzz"))
			go func() { shared.BuzzedInChan <- true }()
		} else {
			log.Println(name, "failed to buzz")
		}
	}
}
