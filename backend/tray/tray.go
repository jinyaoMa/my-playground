package tray

import (
	"github.com/getlantern/systray"
)

var Instance *Tray

type Tray struct {
}

func (t *Tray) onReady() {

}

func (t *Tray) onExit() {

}

func New() *Tray {
	t := &Tray{}
	systray.Register(t.onReady, t.onExit)
	return t
}
