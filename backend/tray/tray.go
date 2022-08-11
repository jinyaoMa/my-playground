package tray

import (
	"embed"
	"encoding/json"
	"my-playground/backend/tray/menus"
	"sync"

	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

//go:embed locales/en.json
//go:embed locales/zh.json
var locales embed.FS

var (
	once sync.Once
	tray *Tray
)

type Tray struct {
	quit menus.Quit
}

func Register() *Tray {
	once.Do(func() {
		tray = &Tray{
			quit: menus.Quit{},
		}
		systray.Run(tray.onReady, tray.onExit)
	})
	return tray
}

func (t *Tray) Quit() {
	systray.Quit()
}

func (t *Tray) onReady() {
	systray.SetIcon(icon)
	systray.SetTooltip("My Playground")
	quitLocale := make(map[string]string)
	rawJson, _ := locales.ReadFile("locales/zh.json")
	json.Unmarshal(rawJson, &quitLocale)
	t.quit.Init(menus.QuitLocale{
		Title:   quitLocale["quit"],
		Tooltip: quitLocale["quit"],
	})
	t.quit.Watch(menus.QuitListener{
		OnQuit: func() {
			systray.Quit()
		},
	})
}

func (t *Tray) onExit() {

}
