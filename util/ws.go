package util

import (
	"encoding/json"
	"time"
	"trivia-app/dlog"

	"github.com/gorilla/websocket"
)

const READ_DEADLINE = 10 // seconds

type WSMsg struct {
	Name    string `json:"name"`
	Headers struct {
		Url string `json:"HX-Current-URL"`
	} `json:"HEADERS"`
}

func HandleWebSocket(conn *websocket.Conn) {
	defer func(conn *websocket.Conn) {
		dlog.DLog("closing websocket connection")
		err := conn.Close()
		if err != nil {
			dlog.DLog("error closing leaderboard websocket")
		}
	}(conn)

	var body WSMsg

	conn.SetReadDeadline(time.Now().Add(3 * READ_DEADLINE * time.Second))
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		err = json.Unmarshal(p, &body)
		if err != nil {
			continue
		}
		msg := body.Name
		dlog.DLog("websocket message:", msg)

		if msg == "\x1F" {
			// ping-pong
			dlog.DLog("keeping websocket alive")
			conn.WriteMessage(websocket.TextMessage, []byte("\x1F"))
			conn.SetReadDeadline(time.Now().Add(3 * READ_DEADLINE * time.Second))
			continue
		}

	}
}
