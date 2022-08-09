package tray

import "github.com/getlantern/systray"

func New() *Tray {
	t := &Tray{}
	systray.Register(t.onReady, t.onExit)
	return t
}
