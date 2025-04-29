package page_handlers

import (
	"net/http"
	env "trivia-app"
	"trivia-app/handlers"
)

type JoinData struct {
	URL    string
	QRCode string
}

func Join(w http.ResponseWriter, r *http.Request) {
	url := "http://" + env.GetIP() + ":8080"

	var qrCodeStr string
	if handlers.QrCodeErr == nil {
		qrCodeStr = handlers.QrCodeObj.ToString(false)
	} else {
		qrCodeStr = ""
	}

	handlers.RenderTemplate(w, "join.html", JoinData{URL: url, QRCode: qrCodeStr})
}
