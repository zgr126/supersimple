package main

import (
	"embed"
	_ "embed"
)

//go:embed admin/dist
var staticFiles embed.FS

func main() {
	run()
}
