package handlers

import (
	"net/http"
	"trivia-app/api/shared"
	"trivia-app/api/util"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	token, err := util.ReadToken(r)
	if err != nil {
		util.InputError(w, util.NO_COOKIE)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		util.InputError(w, util.NO_NAME)
		return
	}

	verified := shared.PlayerStore.VerifyTokenName(token, name)

	if !verified {
		util.InputError(w, util.INVALID_TOKEN)
	} else {
		util.Success(w)
	}
}
