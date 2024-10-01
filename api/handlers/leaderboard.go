package handlers

import (
	"log"
	"net/http"
	"sort"
	"trivia-app/api/shared"
)

var leaderboardWS = shared.NewWebsocketStore()

type leaderboardPlayer struct {
	Name string 
	Score int
}

func Leaderboard(w http.ResponseWriter, r *http.Request) {
	conn, err := shared.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading leaderboard connection")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn.WriteJSON(makeLeaderboard())
	leaderboardWS.InsertConn(conn)

	go leaderboardWS.KeepAlive(conn)
}

func makeLeaderboard() []leaderboardPlayer {
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

	leaderboardList := make([]leaderboardPlayer, len(playerList))
	for i, player := range playerList {
		leaderboardList[i] = leaderboardPlayer{
			Name: player.Name,
			Score: player.Score,
		}
	}

	return leaderboardList
}

func BroadcastLeaderboard() {
	for range shared.LeaderboardChan {
		leaderboard := makeLeaderboard()
		leaderboardWS.WriteToAll(leaderboard)
	}
}
