package main

import (
	"log"
	"os"
	env "trivia-app"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
	"trivia-app/server"
	"trivia-app/shared"
)

func main() {
	os.Setenv("TZ", "America/New_York")

	log.Println("loading .env")
	env.LoadEnv()
	handlers.InitQrCode()
	shared.LoadPassword()

	go ws_handlers.BroadcastLeaderboard()
	go ws_handlers.BroadcastBuzzedIn()
	go ws_handlers.BroadcastPlayerList()

	server := server.Server()
	log.Println("running at " + handlers.QrCodeObj.Content)
	server.ListenAndServe()
}
