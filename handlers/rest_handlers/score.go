package rest_handlers

//
// import (
// 	"net/http"
// 	"strconv"
// 	"trivia-app/shared"
// 	"trivia-app/util"
// )
//
// func Score(w http.ResponseWriter, r *http.Request) {
// 	playerName := r.URL.Query().Get("name")
// 	if playerName == "" {
// 		util.InputError(w, util.NO_NAME)
// 		return
// 	}
//
// 	amountStr := r.URL.Query().Get("amount")
// 	if amountStr == "" {
// 		// util.InputError()
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("no amount provided"))
// 		return
// 	}
// 	amount, err := strconv.ParseInt(amountStr, 10, 64)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("invalid amount"))
// 		return
// 	}
//
// 	token, ok := shared.PlayerStore.NameToToken(playerName)
// 	if !ok {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte("player not found"))
// 		return
// 	}
//
// 	delta := int(amount)
// 	shared.PlayerStore.PutPlayer(token, shared.UpdatePlayer{ScoreDiff: &delta})
//
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("success"))
// }

const N = "no"
