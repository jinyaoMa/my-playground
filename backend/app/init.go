package app

import "my-playground/backend/tray"

var app *App

func Lication() *App {
	if app != nil {
		return app
	}

	return &App{
		tray: &tray.Tray{},
	}
}
