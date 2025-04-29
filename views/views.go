package views

import "embed"

//go:embed layout.html
var Layout string

//go:embed pages
var Pages embed.FS

//go:embed components
var Components embed.FS
