package main

import (
	"embed"
	"my-playground/backend/panel"
)

//go:embed frontend/panel/dist
var FnPanel embed.FS

func main() {
	panel.Run(FnPanel)
}
