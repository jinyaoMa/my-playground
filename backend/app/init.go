package app

import (
	"my-playground/backend/database"
	"my-playground/backend/tray"
)

var app *App

func Lication() *App {
	if app != nil {
		return app
	}

	return &App{
		Tray: tray.New(),
		DB:   database.New(),
	}
}
