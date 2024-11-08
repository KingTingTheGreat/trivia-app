package handlers

import (
	"net/http"
	"time"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
	"trivia-app/api/util"

	"github.com/gorilla/websocket"
)

const READ_DEADLINE = 10

func BuzzWs(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("web socket")
	// read cookie
	token, err := util.ReadToken(r)
	if err != nil {
		util.UserInputError(w, "no cookie")
		return
	}

	dlog.DLog(token)

	// player associated with cookie
	player, ok := shared.PlayerStore.GetPlayer(token)
	if !ok {
		dlog.DLog("player does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dlog.DLog("player exists")

	dlog.DLog("upgrading to websocket")
	// upgrade to websocket
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading connection to websocket")
		return
	}

	dlog.DLog("handling created websocket")
	go websocketHandler(conn, token, player)
}

func websocketHandler(conn *websocket.Conn, token string, player *shared.Player) {
	dlog.DLog("websocketHandler()")
	defer func(conn *websocket.Conn) {
		dlog.DLog("closing websocket", player.Name)
		shared.PlayerStore.NilPlayerWS(token)
		err := conn.Close()
		if err != nil {
			dlog.DLog("error closing websocket connection")
		}
	}(conn)

	dlog.DLog("reading message")
	// first message is used to verify player, not buzz in
	_, p, err := conn.ReadMessage()
	if err != nil {
		dlog.DLog("could not read from websocket")
		return
	}
	name := string(p)
	dlog.DLog("websocket", name)
	if player.Name != name {
		dlog.DLog("name and token do not match")
		return
	}

	// update player store entry with websocket
	shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{Websocket: conn})

	// update client-side with correct button state
	if player.ButtonReady {
		dlog.DLog(player.Name, "ready")
		conn.WriteMessage(websocket.TextMessage, []byte("ready"))
	} else {
		dlog.DLog(player.Name, "buzz")
		conn.WriteMessage(websocket.TextMessage, []byte("buzz"))
	}

	go func() { shared.LeaderboardChan <- true }()

	killChan := make(chan bool)
	readFunc := func() {
		defer func() { dlog.DLog("ending readFunc", player.Name) }()
		conn.SetReadDeadline(time.Now().Add(READ_DEADLINE * time.Second))
		for {
			dlog.DLog("buzz ws loop")

			// await message from client
			_, p, err := conn.ReadMessage()
			if err != nil {
				dlog.DLog("websocket reading error", player.Name, err.Error())
				return
			}
			msg := string(p)
			dlog.DLog("websocket message:", name, msg)

			select {
			case <-killChan:
				return
			default:
				if msg == "\x1F" {
					// ping-pong
					dlog.DLog(name, "keeping websocket alive")
					conn.SetReadDeadline(time.Now().Add(READ_DEADLINE * time.Second))
				} else if msg == name && shared.PlayerStore.BuzzIn(token, msg) {
					// buzz in
					dlog.DLog(name, "buzzed in")
					conn.WriteMessage(websocket.TextMessage, []byte("buzz"))
					go func() { shared.BuzzedInChan <- true }()
				} else {
					dlog.DLog(name, "failed to buzz")
				}
			}
		}
	}

	go readFunc()

	select {
	case b := <-player.WsClose:
		killChan <- b
	case <-killChan:
		return
	}
}
