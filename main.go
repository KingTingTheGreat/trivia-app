package main

import (
	"log"
	"trivia-app/api/handlers"
	"trivia-app/api/server"
	"trivia-app/api/shared"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	shared.LoadPassword()

	go handlers.BroadcastLeaderboard()
	go handlers.BroadcastBuzzedIn()
	go handlers.BroadcastPlayerList()

	server := server.Server()
	log.Println("api running at http://localhost:8080")
	server.ListenAndServe()
}
