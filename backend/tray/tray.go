package tray

import "github.com/getlantern/systray"

type Tray struct {
}

func (t *Tray) Quit() {
	systray.Quit()
}

func (t *Tray) onReady() {

}

func (t *Tray) onExit() {

}
