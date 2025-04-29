package shared

import (
	env "trivia-app"
)

func ReactiveBuzzers() bool {
	return env.EnvVal("REACTIVE") != "false"
}
