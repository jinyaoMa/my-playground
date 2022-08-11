package app

import (
	"context"
	"sync"
)

var (
	once sync.Once
	app  *App
)

// App struct
type App struct {
	ctx      context.Context // bind wails backend
	database *Database
	tray     *Tray
	server   *Server
}

func Lication(ctx ...context.Context) *App {
	once.Do(func() {
		app = &App{
			database: NewDatabase(),
			tray:     NewTray(),
			server:   NewServer(),
		}
	})
	if len(ctx) > 0 {
		app.ctx = ctx[0]
	}
	return app
}

func (a *App) QuitTray() {
	a.tray.quit.Trigger()
}
