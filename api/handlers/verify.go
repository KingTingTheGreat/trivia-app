package handlers

import (
	"net/http"
	"trivia-app/api/shared"
	"trivia-app/api/util"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	token, err := util.ReadToken(r)
	if err != nil {
		util.UserInputError(w, "no cookie")
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		util.UserInputError(w, "no name")
		return
	}

	verified := shared.PlayerStore.VerifyTokenName(token, name)

	if !verified {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("failure"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}
}
