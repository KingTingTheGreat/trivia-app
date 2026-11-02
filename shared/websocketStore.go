package shared

import (
	"sync"
	"trivia-app/dlog"
	"trivia-app/util"

	"github.com/gorilla/websocket"
)

type websocketStore struct {
	mu            sync.RWMutex
	websocketData map[*websocket.Conn]bool
}

func (ws *websocketStore) InsertConn(conn *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.websocketData[conn] = true
}

func (ws *websocketStore) DeleteConn(conn *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	delete(ws.websocketData, conn)
}

func (ws *websocketStore) WriteToAll(data []byte) {
	dlog.DLog("write to all")
	ws.mu.Lock()
	defer ws.mu.Unlock()
	dlog.DLog("got lock", len(ws.websocketData))

	for conn := range ws.websocketData {
		dlog.DLog("writing to conn")
		conn.WriteMessage(websocket.TextMessage, data)
	}
}

func (ws *websocketStore) KeepAlive(conn *websocket.Conn) {
	defer func(conn *websocket.Conn) {
		// closing websocket is handled in util.HandleWebSocket
		dlog.DLog("deleting conn from websocket store")
		ws.DeleteConn(conn)
	}(conn)
	util.HandleWebSocket(conn)
}

func NewWebsocketStore() websocketStore {
	return websocketStore{
		mu:            sync.RWMutex{},
		websocketData: make(map[*websocket.Conn]bool),
	}
}
