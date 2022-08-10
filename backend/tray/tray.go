package tray

import (
	_ "embed"
	"sync"

	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

var (
	once sync.Once
	tray *Tray
)

type Tray struct {
}

func Register() *Tray {
	once.Do(func() {
		tray = &Tray{}
		systray.Register(tray.onReady, tray.onExit)
	})
	return tray
}

func (t *Tray) Quit() {
	systray.Quit()
}

func (t *Tray) onReady() {
	systray.SetIcon(icon)

}

func (t *Tray) onExit() {

}
