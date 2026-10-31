package util

import "math/rand"

var colors = []string{
	"bg-fuchsia-800",
	"bg-green-600",
	"bg-blue-600",
	"bg-rose-600",
	"bg-orange-500",
	"bg-cyan-400",
	"bg-purple-500",
}

func RandomTwColor() string {
	return colors[rand.Intn(len(colors))]
}
