package main

import (
	"log"
	env "trivia-app"
	"trivia-app/handlers/ws_handlers"
	"trivia-app/server"
	"trivia-app/shared"
)

func main() {
	env.LoadEnv()
	shared.LoadPassword()

	go ws_handlers.BroadcastLeaderboard()
	go ws_handlers.BroadcastBuzzedIn()
	go ws_handlers.BroadcastPlayerList()

	server := server.Server()
	log.Println("running at http://" + env.GetIP() + ":8080")
	server.ListenAndServe()
}
