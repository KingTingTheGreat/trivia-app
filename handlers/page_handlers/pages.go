package page_handlers

import (
	"net/http"
	"trivia-app/handlers"
)

type PageSection struct {
	Title       string
	Description string
	Path        string
}

type PagesData struct {
	Pages []PageSection
}

func Pages(w http.ResponseWriter, r *http.Request) {
	data := PagesData{
		Pages: []PageSection{
			{
				Title:       "Home",
				Description: "this is where players join the game",
				Path:        "/",
			},
			{
				Title:       "Join",
				Description: "Display a QR Code players can scan to access the application",
				Path:        "/join",
			},
			{
				Title:       "Leaderboard",
				Description: "see players ranked by score. ties are broken by who reached that score first.",
				Path:        "/leaderboard",
			},
			{
				Title:       "Buzzed In",
				Description: "view players as they buzz in",
				Path:        "/buzzed-in",
			},
			{
				Title:       "Game Info",
				Description: "buzzed in and leaderboard",
				Path:        "/game",
			},
			{
				Title:       "Control",
				Description: "access game controls",
				Path:        "/control",
			},
		},
	}

	handlers.RenderTemplate(w, "pages.html", data)
}
