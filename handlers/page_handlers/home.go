package page_handlers

import (
	"net/http"
	"trivia-app/handlers"
	"trivia-app/shared"
)

type HomeData struct {
	Name  string
	Error string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := HomeData{}
	ctx := r.Context()
	player, ok := ctx.Value("player").(*shared.Player)
	if ok {
		data.Name = player.Name
	}

	handlers.RenderTemplate(w, "home.html", data)
}
