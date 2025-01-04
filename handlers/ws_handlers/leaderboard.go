package ws_handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
	"time"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/shared"

	"github.com/gorilla/websocket"
)

var leaderboardWS = shared.NewWebsocketStore()

type leaderboardPlayer struct {
	Name  string
	Score string
}

type LeaderboardData struct {
	TableId  string
	Title    string
	Headers  []string
	RowData  [][]string
	Endpoint string
}

func LeaderboardWS(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("leaderboard WS")
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading leaderboard connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	leaderboardWS.InsertConn(conn)

	// write current state (this is for reconnections)
	var buf bytes.Buffer
	handlers.RenderComponent(&buf, "table-body.html", LeaderboardData{
		TableId: "leaderboard",
		Title:   "Leaderboard",
		Headers: []string{
			"Name",
			"Score",
		},
		RowData: MakeLeaderboard(),
	})
	conn.WriteMessage(websocket.TextMessage, buf.Bytes())

	go leaderboardWS.KeepAlive(conn)
}

func MakeLeaderboard() [][]string {
	playerList := shared.PlayerStore.AllPlayers()

	// sort by score, then last updated, then name
	sort.Slice(playerList, func(i, j int) bool {
		pI, pJ := playerList[i], playerList[j]
		if pI.Score != pJ.Score {
			return pI.Score > pJ.Score
		}
		if pI.LastUpdate != pJ.LastUpdate {
			return pI.LastUpdate.Before(pJ.LastUpdate)
		}
		return pI.Name < pJ.Name
	})

	leaderboardList := make([][]string, len(playerList))
	for i, player := range playerList {
		leaderboardList[i] = []string{
			player.Name,
			fmt.Sprintf("%d", player.Score),
		}
	}

	return leaderboardList
}

func BroadcastLeaderboard() {
	var buf bytes.Buffer
	for range shared.LeaderboardChan {
		now := time.Now()
		for time.Since(now) < 50*time.Millisecond {
			select {
			case <-shared.LeaderboardChan:
			default:
				break
			}
		}
		handlers.RenderComponent(&buf, "table-body.html", LeaderboardData{
			TableId: "leaderboard",
			Title:   "Leaderboard",
			Headers: []string{
				"Name",
				"Score",
			},
			RowData: MakeLeaderboard(),
		})
		leaderboardWS.WriteToAll(buf.Bytes())
	}
}
