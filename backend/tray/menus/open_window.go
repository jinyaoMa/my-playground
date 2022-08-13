package menus

import (
	"github.com/getlantern/systray"
)

type OpenWindowListener struct {
	OnOpenWindow func()
}

type OpenWindow struct {
	item     *systray.MenuItem
	chanStop chan struct{}
}

func NewOpenWindow() *OpenWindow {
	return &OpenWindow{
		item:     systray.AddMenuItem("", ""),
		chanStop: make(chan struct{}, 1),
	}
}

func (ow *OpenWindow) SetIcon(icon []byte) *OpenWindow {
	ow.item.SetIcon(icon)
	return ow
}

func (ow *OpenWindow) SetLocale(locale map[string]string) *OpenWindow {
	ow.item.SetTitle(locale["openWindow"])
	ow.item.SetTooltip(locale["openWindow"])
	return ow
}

func (ow *OpenWindow) Watch(listener OpenWindowListener) *OpenWindow {
	go func() {
		for {
			select {
			case <-ow.item.ClickedCh:
				listener.OnOpenWindow()
			case <-ow.chanStop:
				return
			}
		}
	}()
	return ow
}

func (ow *OpenWindow) StopWatch() *OpenWindow {
	close(ow.chanStop)
	return ow
}

func (ow *OpenWindow) Click() *OpenWindow {
	ow.item.ClickedCh <- struct{}{}
	return ow
}
