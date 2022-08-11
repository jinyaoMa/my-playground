package app

import (
	"embed"
	"encoding/json"
	"my-playground/backend/menus"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed static/icon.ico
var icon []byte

//go:embed locales/en.json
//go:embed locales/zh.json
var locales embed.FS

type Tray struct {
	quit *menus.Quit
}

func NewTray() *Tray {
	t := &Tray{}
	go systray.Run(t.onReady, t.onExit)
	return t
}

func (t *Tray) onReady() {
	systray.SetIcon(icon)
	systray.SetTooltip("My Playground")

	quitLocale := make(map[string]string)
	rawJson, _ := locales.ReadFile("locales/zh.json")
	json.Unmarshal(rawJson, &quitLocale)

	t.quit = menus.NewQuit(menus.QuitLocale{
		Title:   quitLocale["quit"],
		Tooltip: quitLocale["quit"],
	}).Watch(menus.QuitListener{
		OnQuit: func() {
			systray.Quit()
		},
	})
}

func (t *Tray) onExit() {
	runtime.Quit(app.ctx)
}
