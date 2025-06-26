package env

import (
	_ "embed"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//go:embed .env
var envString string
var env map[string]string

func LoadEnv() {
	var err error
	env, err = godotenv.Unmarshal(envString)
	if err != nil {
		log.Fatal("failed to parse environment variables in .env")
	}
}

func EnvVal(key string) string {
	val := env[key]
	if val != "" {
		return val
	}
	return os.Getenv(key)
}

func GetIP() string {
	ip := EnvVal("IP")
	if ip == "" {
		return "localhost"
	}
	return ip
}

func Set(key string, value string) {
	env[key] = value
}

func Save() error {
	return godotenv.Write(env, ".env")
}
