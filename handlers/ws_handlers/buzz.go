package ws_handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/shared"
	"trivia-app/util"

	"github.com/gorilla/websocket"
)

const READ_DEADLINE = 10 // seconds

type WSMsg struct {
	Name    string `json:"name"`
	Headers struct {
		Url string `json:"HX-Current-URL"`
	} `json:"HEADERS"`
}

func parseName(urlString string) string {
	urlSplit := strings.Split(urlString, "/")
	name, err := url.QueryUnescape(urlSplit[len(urlSplit)-1])
	if err != nil {
		return ""
	}
	return name
}

func BuzzWS(w http.ResponseWriter, r *http.Request) {
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
		util.RedirectError(w, r, "player does not exist")
		return
	}
	dlog.DLog("player exists")

	dlog.DLog("upgrading to websocket")
	// upgrade to websocket
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading connection to websocket")
		util.RedirectError(w, r, "something went wrong, please try again.")
		return
	}

	dlog.DLog("handling created websocket")
	websocketHandler(conn, token, &player)
}

func websocketHandler(conn *websocket.Conn, token string, playerRef *shared.Player) {
	player := *playerRef
	dlog.DLog("websocketHandler()")
	defer func(conn *websocket.Conn) {
		dlog.DLog("closing websocket", player.Name)
		shared.PlayerStore.NilPlayerWS(token, conn)
		err := conn.Close()
		if err != nil {
			dlog.DLog("error closing websocket connection")
		}
	}(conn)
	var body WSMsg

	dlog.DLog("reading message")
	// first message is used to verify player, not buzz in
	_, p, err := conn.ReadMessage()
	if err != nil {
		dlog.DLog("could not read from websocket")
		return
	}

	err = json.Unmarshal(p, &body)
	if err != nil {
		dlog.DLog("error decoding message from websocket", string(p))
		return
	}

	name := parseName(body.Headers.Url)
	if player.Name != name {
		dlog.DLog("name and token do not match", name, token)
		return
	}

	// update player store entry with websocket
	player, err = shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{Websocket: conn})
	if err != nil {
		dlog.DLog("race condition!!!")
		return
	}

	var buf bytes.Buffer
	handlers.RenderComponent(&buf, "buzz-button.html", handlers.Play{Ready: player.ButtonReady})
	conn.WriteMessage(websocket.TextMessage, buf.Bytes())

	go func() { shared.LeaderboardChan <- true }()

	killChan := make(chan bool)
	readFunc := func() {
		defer func() { dlog.DLog("ending readFunc", player.Name); killChan <- true }()
		conn.SetReadDeadline(time.Now().Add(READ_DEADLINE * time.Second))
		for {
			dlog.DLog("buzz ws loop")

			// await message from client
			_, p, err := conn.ReadMessage()
			if err != nil {
				dlog.DLog("websocket reading error: ", player.Name, err.Error())
				return
			}
			err = json.Unmarshal(p, &body)
			if err != nil {
				dlog.DLog("error unmarshaling from ws", player.Name, err.Error())
				// return
				continue
			}
			msg := body.Name
			dlog.DLog("websocket message:", name, msg)

			if msg == "\x1F" {
				// ping-pong
				dlog.DLog(name, "keeping websocket alive")
				conn.SetReadDeadline(time.Now().Add(READ_DEADLINE * time.Second))
			} else if msg == name && shared.PlayerStore.BuzzIn(token, msg) {
				// buzz in
				dlog.DLog(name, "buzzed in")
				handlers.RenderComponent(&buf, "buzz-button.html", handlers.Play{Ready: false})
				conn.WriteMessage(websocket.TextMessage, buf.Bytes())

				go func() { shared.BuzzedInChan <- true }()
			} else {
				dlog.DLog(name, "failed to buzz")
			}
		}
	}

	go readFunc()

	select {
	case <-player.WsClose:
		dlog.DLog("wsclose chan")
		return
	case <-killChan:
		dlog.DLog("kill chan")
		return
	}
}
