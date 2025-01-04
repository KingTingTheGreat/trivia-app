package page_handlers

import (
	"log"
	"net/http"
	"trivia-app/handlers"
	"trivia-app/shared"
	"trivia-app/util"
)

type PlayData struct {
	Name  string
	Ready bool
}

func Play(w http.ResponseWriter, r *http.Request) {
	log.Println("play handler")
	name := r.PathValue("name")

	ctx := r.Context()
	player, ok := ctx.Value("player").(*shared.Player)
	if !ok || player.Name != name {
		log.Println("player does not exist")
		util.RedirectError(w, r, "")
		return
	}

	data := PlayData{
		Name:  name,
		Ready: player.ButtonReady,
	}

	log.Println("name", name)
	handlers.RenderTemplate(w, "play.html", data)
}
