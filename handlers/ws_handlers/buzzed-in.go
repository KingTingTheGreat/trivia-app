package ws_handlers

import (
	"bytes"
	"net/http"
	"sort"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/shared"

	"github.com/gorilla/websocket"
)

var buzzedInWS = shared.NewWebsocketStore()

type buzzedInPlayer struct {
	Name string
	Time string
}

func BuzzedInWS(w http.ResponseWriter, r *http.Request) {
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading buzzed in connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buzzedInWS.InsertConn(conn)

	// write current state (this is for reconnections)
	var buf bytes.Buffer
	handlers.RenderComponent(&buf, "table-body.html", LeaderboardData{
		TableId: "buzzed-in",
		Title:   "Buzzed In",
		Headers: []string{
			"Name",
			"Time",
		},
		RowData: MakeBuzzedIn(),
	})
	conn.WriteMessage(websocket.TextMessage, buf.Bytes())

	go buzzedInWS.KeepAlive(conn)
}

func MakeBuzzedIn() [][]string {
	playerList := shared.PlayerStore.AllPlayers()

	// sort by buzz in time, then score, then name
	sort.Slice(playerList, func(i, j int) bool {
		pI, pJ := playerList[i], playerList[j]
		if pI.BuzzedIn != pJ.BuzzedIn {
			return pI.BuzzedIn.Before(pJ.BuzzedIn)
		}
		if pI.Score != pJ.Score {
			return pI.Score > pJ.Score
		}
		return pI.Name < pJ.Name
	})

	buzzedIn := make([][]string, 0)
	for _, player := range playerList {
		// filter out players who have not buzzed in
		if player.BuzzedIn.IsZero() {
			// dlog.DLog("skiping buzz")
			continue // could probably break loop instead
		}
		dlog.DLog("appending buzz")
		buzzedIn = append(buzzedIn, []string{
			player.Name,
			player.BuzzedIn.Format("03:04:05.0 PM"),
		})
	}

	return buzzedIn
}

func BroadcastBuzzedIn() {
	var buf bytes.Buffer
	for range shared.BuzzedInChan {
		dlog.DLog("buzzed in chan")
		handlers.RenderComponent(&buf, "table-body.html", LeaderboardData{
			TableId: "buzzed-in",
			Title:   "Buzzed In",
			Headers: []string{
				"Name",
				"Time",
			},
			RowData: MakeBuzzedIn(),
		})
		buzzedInWS.WriteToAll(buf.Bytes())
	}
}
