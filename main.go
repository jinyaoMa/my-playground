package main

import (
	"embed"
	"my-services/backend/window"
)

//go:embed frontend/app/dist
var frontend embed.FS

func main() {
	window.Run(frontend)
}
