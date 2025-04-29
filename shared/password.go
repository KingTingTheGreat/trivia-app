package shared

import (
	env "trivia-app"
)

var Password string

// assumes env as been loaded
func LoadPassword() {
	Password = env.EnvVal("PASSWORD")

	if Password == "" {
		panic("No password found in .env file")
	}
}
