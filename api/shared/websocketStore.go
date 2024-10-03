package shared

import (
	"log"
	"sync"

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

func (ws *websocketStore) WriteToAll(data interface{}) {
	log.Println("write to all")
	ws.mu.Lock()
	defer ws.mu.Unlock()
	log.Println("got lock", len(ws.websocketData))

	for conn := range ws.websocketData {
		log.Println("writing to conn")
		go conn.WriteJSON(data)
	}
}

func (ws *websocketStore) KeepAlive(conn *websocket.Conn) {
	defer func(conn *websocket.Conn) {
		log.Println("closing leaderboard websocket")
		ws.DeleteConn(conn)
		err := conn.Close()
		if err != nil {
			log.Println("error closing leaderboard websocket")
		}
	}(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func NewWebsocketStore() websocketStore {
	return websocketStore{
		mu:            sync.RWMutex{},
		websocketData: make(map[*websocket.Conn]bool),
	}
}
