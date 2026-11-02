package ws_handlers

import (
	"bytes"
	"fmt"
	"io"
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
	TableId      string
	Title        string
	Description  string
	EmptyMessage string
	TableData    []TableRow
	Endpoint     string
}

func rankColor(place int) string {
	switch place {
	case 1:
		return "#FFD700"
	case 2:
		return "#C0C0C0"
	case 3:
		return "#CD7F32"
	default:
		return ""
	}
}

func renderLeaderboard(w io.Writer) {
	handlers.RenderComponent(w, "table-body.html", LeaderboardData{
		TableId:      "leaderboard",
		Title:        "Leaderboard",
		Description:  "Top players by score",
		EmptyMessage: "Players will appear here once they join",
		TableData:    MakeLeaderboard(),
	})

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
	renderLeaderboard(&buf)
	conn.WriteMessage(websocket.TextMessage, buf.Bytes())

	go leaderboardWS.KeepAlive(conn)
}

func MakeLeaderboard() []TableRow {
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

	leaderboardList := make([]TableRow, len(playerList))
	for i, player := range playerList {
		place := i + 1
		icon := false
		if place <= 3 {
			icon = true
		}
		leaderboardList[i] = TableRow{
			Place:    place,
			Name:     player.Name,
			Data:     fmt.Sprintf("%d", player.Score),
			DataDesc: "Points",
			Icon:     icon,
			Color:    rankColor(place),
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
		renderLeaderboard(&buf)
		leaderboardWS.WriteToAll(buf.Bytes())
	}
}
