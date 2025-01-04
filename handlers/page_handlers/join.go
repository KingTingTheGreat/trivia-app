package page_handlers

import (
	"log"
	"net/http"
	"os"
	"trivia-app/handlers"

	qrcode "github.com/skip2/go-qrcode"
)

type JoinData struct {
	URL    string
	QRCode string
}

func Join(w http.ResponseWriter, r *http.Request) {
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "localhost"
	}

	url := "http://" + ip + ":8080"

	qrCode, err := qrcode.New(url, qrcode.High)
	if err != nil {
		log.Println(err)
	}

	log.Println(qrCode.WriteFile(256, "./public/qrcode.png"))
	handlers.RenderTemplate(w, "join.html", JoinData{URL: url, QRCode: qrCode.ToString(false)})
}
