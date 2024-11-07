package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const envFile = ".env.local"

func main() {
	godotenv.Load(envFile)

	envVals := map[string]string{
		"PASSWORD": "",
		"IP":       "",
	}

	password := os.Getenv("PASSWORD")
	for password == "" {
		fmt.Print("Input a password: ")
		_, err := fmt.Scanln(&password)
		if err == nil {
			password = strings.TrimSpace(password)
			if password != "" {
				break
			}
		}
	}
	envVals["PASSWORD"] = password

	ip := os.Getenv("IP")
	res, err := http.Get("https://ifconfig.me/ip")
	if err == nil {
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err == nil {
			ip = string(body)
		}
	}
	envVals["IP"] = ip

	godotenv.Write(envVals, envFile)
}
