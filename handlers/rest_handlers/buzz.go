package rest_handlers

import (
	"net/http"
	"trivia-app/dlog"
	"trivia-app/shared"
	"trivia-app/util"
)

func Buzz(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("buzz rest handler")
	ctx := r.Context()
	token, ok := ctx.Value("token").(string)
	if !ok {
		dlog.DLog("no token")
		util.RedirectError(w, r, "something went wrong, please try again.")
		return
	}

	name := util.ReadValue(r, "name")

	success, err := shared.PlayerStore.BuzzIn(token, name)
	if err != nil {
		dlog.DLog("buzz in error", err.Error())
		util.RedirectError(w, r, "something went wrong, please try again.")
		return
	} else if success {
		go func() { shared.BuzzedInChan <- true }()
	}

	w.WriteHeader(http.StatusOK)
	if success {
		dlog.DLog("successfully buzzed in")
		w.Write([]byte("success"))
	} else {
		dlog.DLog("failed to buzz in")
		w.Write([]byte("already buzzed in"))
	}
}
