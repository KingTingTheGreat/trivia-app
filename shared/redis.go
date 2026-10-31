package shared

// import (
// 	"bytes"
// 	"context"
// 	"crypto/tls"
// 	"encoding/gob"
// 	"fmt"
// 	"log"
// 	"time"
// 	env "trivia-app"
//
// 	"github.com/redis/go-redis/v9"
// )
//
// func RedisClient() (*redis.Client, error) {
// 	redisHost := env.EnvVal("REDIS_PRIMARY_ENDPOINT")
// 	redisPort := env.EnvVal("REDIS_PORT")
// 	redisPassword := env.EnvVal("REDIS_PASSWORD")
//
// 	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
// 	log.Println("redis addr", addr)
//
// 	redisOpts := redis.Options{
// 		Addr:      addr,
// 		Password:  redisPassword,
// 		TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
// 	}
//
// 	rdb := redis.NewClient(&redisOpts)
//
// 	_, err := rdb.Ping(context.Background()).Result()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to redis")
// 	}
//
// 	return rdb, nil
// }
//
// func TestRedis() {
// 	ctx := context.Background()
//
// 	player := Player{
// 		CleanName:        "clean name",
// 		Name:             "name",
// 		Score:            0,
// 		ButtonReady:      true,
// 		CorrectQuestions: make([]string, 0),
// 		LastUpdate:       time.Now(),
// 		BuzzedIn:         time.Time{},
// 		Websocket:        nil,
// 		WsClose:          make(chan bool, 1),
// 	}
//
// 	key := "test"
//
// 	var buf bytes.Buffer
// 	err := gob.NewEncoder(&buf).Encode(player)
// 	if err != nil {
// 		log.Println("failed to encode player", err)
// 		return
// 	}
//
// 	rdb, _ := RedisClient()
// 	err = rdb.Set(ctx, key, buf.Bytes(), 0).Err()
// 	if err != nil {
// 		log.Println("failed to insert redis key", err)
// 		return
// 	}
//
// 	byte, err := rdb.Get(ctx, key).Bytes()
// 	if err != nil {
// 		log.Println("failed to get value from redis", err)
// 		return
// 	}
//
// 	b := bytes.NewReader(byte)
//
// 	var res Player
//
// 	err = gob.NewDecoder(b).Decode(&res)
// 	if err != nil {
// 		log.Println("failed to decode player", err)
// 		return
// 	}
//
// 	fmt.Println("player", player)
// 	fmt.Println("res", res)
// }
