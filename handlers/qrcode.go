package handlers

import (
	"net/http"
	env "trivia-app"

	"github.com/skip2/go-qrcode"
)

var QrCodeObj *qrcode.QRCode
var QrCodeErr error

// InitQrCode must be called after env.LoadEnv so the IP from .env is used.
func InitQrCode() {
	QrCodeObj, QrCodeErr = qrcode.New(env.GetIP(), qrcode.High)
}

func QrCode(w http.ResponseWriter, r *http.Request) {
	if QrCodeErr == nil {
		QrCodeObj.Write(256, w)
	}
}
