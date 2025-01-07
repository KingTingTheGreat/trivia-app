package shared

import "os"

func ReactiveBuzzers() bool {
	return os.Getenv("REACTIVE") != "false"
}
