package app

import (
	"my-playground/backend/database"
	"my-playground/backend/tray"
	"sync"
)

var once sync.Once
var app *App

func Lication() *App {
	once.Do(func() {
		app = &App{
			Tray: tray.New(),
			DB:   database.New(),
		}
	})

	return app
}
