package backend

import (
	"embed"
	"my-playground/backend/menus"
	"my-playground/backend/utils"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed .assets/favicon.ico
var favicon []byte

//go:embed locales/en.json
//go:embed locales/zh.json
var locales embed.FS

var tray *Tray

type Tray struct {
	openWindow *menus.OpenWindow
	quit       *menus.Quit
}

func setupTray() {
	tray = &Tray{}
	go systray.Run(tray.onReady, tray.onExit)
}

func (t *Tray) onReady() {
	systray.SetIcon(favicon)

	rawJson, _ := locales.ReadFile("locales/zh.json")
	locale := utils.GetLocaleFromJSON(rawJson)

	systray.SetTooltip(locale["appname"])

	t.openWindow = menus.
		NewOpenWindow().
		SetLocale(locale).
		Watch(menus.OpenWindowListener{
			OnOpenWindow: func() {
				runtime.WindowShow(app.ctx)
			},
		})

	systray.AddSeparator()

	t.quit = menus.
		NewQuit().
		SetLocale(locale).
		Watch(menus.QuitListener{
			OnQuit: func() {
				systray.Quit()
			},
		})
}

func (t *Tray) onExit() {
	t.openWindow.StopWatch()
	t.quit.StopWatch()
	runtime.Quit(app.ctx)
}
