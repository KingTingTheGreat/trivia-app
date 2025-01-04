package shared

import (
	"os"
)

var Password string

// assumes env as been loaded
func LoadPassword() {
	Password = os.Getenv("PASSWORD")

	if Password == "" {
		panic("No password found in .env file")
	}
}
