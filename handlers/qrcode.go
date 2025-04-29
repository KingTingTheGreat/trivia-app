package handlers

import (
	"net/http"
	env "trivia-app"

	"github.com/skip2/go-qrcode"
)

var QrCodeObj, QrCodeErr = qrcode.New("http://"+env.GetIP()+":8080", qrcode.High)

func QrCode(w http.ResponseWriter, r *http.Request) {
	if QrCodeErr == nil {
		QrCodeObj.Write(256, w)
	}
}
