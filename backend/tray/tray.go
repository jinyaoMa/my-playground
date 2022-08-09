package tray

import (
	_ "embed"

	"github.com/getlantern/systray"
)

//go:embed icons/icon.ico
var icon []byte

type Tray struct {
}

func (t *Tray) Quit() {
	systray.Quit()
}

func (t *Tray) onReady() {
	systray.SetIcon(icon)

}

func (t *Tray) onExit() {

}
