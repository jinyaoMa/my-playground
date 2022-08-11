package app

import (
	"my-playground/backend/server"
	"my-playground/backend/tray"
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	app  *App
)

type App struct {
	DB     *gorm.DB
	Tray   *tray.Tray
	Server *server.Server
}

func Lication() *App {
	once.Do(func() {
		db := connectDB()
		t := tray.Register()
		s := server.Load()

		app = &App{
			DB:     db,
			Tray:   t,
			Server: s,
		}
	})
	return app
}
