package main

import (
	"log"
	env "trivia-app"
	"trivia-app/handlers/ws_handlers"
	"trivia-app/server"
	"trivia-app/shared"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	shared.LoadPassword()

	go ws_handlers.BroadcastLeaderboard()
	go ws_handlers.BroadcastBuzzedIn()
	go ws_handlers.BroadcastPlayerList()

	server := server.Server()
	log.Println("running at http://" + env.EnvVal("IP") + ":8080")
	server.ListenAndServe()
}
