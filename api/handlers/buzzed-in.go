package handlers

import (
	"net/http"
	"sort"
	"trivia-app/api/dlog"
	"trivia-app/api/shared"
)

var buzzedInWS = shared.NewWebsocketStore()

type buzzedInPlayer struct {
	Name string
	Time string
}

func BuzzedIn(w http.ResponseWriter, r *http.Request) {
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		dlog.DLog("error upgrading buzzed in connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn.WriteJSON(makeBuzzedIn())
	buzzedInWS.InsertConn(conn)

	go buzzedInWS.KeepAlive(conn)
}

func makeBuzzedIn() []buzzedInPlayer {
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

	buzzedIn := make([]buzzedInPlayer, 0)
	for _, player := range playerList {
		// filter out players who have not buzzed in
		if player.BuzzedIn.IsZero() {
			// dlog.DLog("skiping buzz")
			continue // could probably break loop instead
		}
		dlog.DLog("writing buzz")
		buzzedIn = append(buzzedIn, buzzedInPlayer{
			Name: player.Name,
			Time: player.BuzzedIn.Format("03:04:05.0 PM"),
		})
	}

	// dlog.DLog(buzzedIn)

	return buzzedIn
}

func BroadcastBuzzedIn() {
	for range shared.BuzzedInChan {
		dlog.DLog("buzzed in chan")
		buzzedIn := makeBuzzedIn()
		buzzedInWS.WriteToAll(buzzedIn)
	}
}
