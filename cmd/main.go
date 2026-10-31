package main

import (
	// "context"
	// "crypto/tls"
	// "fmt"
	"log"
	"os"
	env "trivia-app"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
	"trivia-app/server"
	"trivia-app/shared"
	// "github.com/redis/go-redis/v9"
)

// func testRedis() error {
// 	log.Println("getting redis environment variables")
// 	redisHost := env.EnvVal("REDIS_PRIMARY_ENDPOINT")
// 	redisPort := env.EnvVal("REDIS_PORT")
// 	redisPassword := env.EnvVal("REDIS_PASSWORD")
//
// 	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
// 	log.Println("redis addr", addr)
//
// 	redisOpts := redis.Options{
// 		Addr:     addr,
// 		Password: redisPassword,
// 	}
//
// 	_, ok := os.LookupEnv("CAPTAIN_APP_NAME")
// 	if ok {
// 		redisOpts.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
// 	}
//
// 	rdb := redis.NewClient(&redisOpts)
// 	defer rdb.Close()
//
// 	ctx := context.Background()
//
// 	pong, err := rdb.Ping(ctx).Result()
// 	if err != nil {
// 		log.Println("Could not connect to Redis:", err)
// 		return err
// 	}
//
// 	fmt.Println("Connected to Redis!", pong)
// 	return nil
// }

func main() {
	os.Setenv("TZ", "America/New_York")

	log.Println("loading .env")
	env.LoadEnv()
	shared.LoadPassword()

	// err := testRedis()
	// if err != nil {
	// 	log.Println("redis test failed", err)
	// }
	//
	// shared.TestRedis()

	go ws_handlers.BroadcastLeaderboard()
	go ws_handlers.BroadcastBuzzedIn()
	go ws_handlers.BroadcastPlayerList()

	server := server.Server()
	log.Println("running at " + handlers.QrCodeObj.Content)
	server.ListenAndServe()
}
