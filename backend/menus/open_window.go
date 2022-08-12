package menus

import (
	"github.com/getlantern/systray"
)

type OpenWindowListener struct {
	OnOpenWindow func()
}

type OpenWindow struct {
	item           *systray.MenuItem
	chanOpenWindow chan struct{}
}

func NewOpenWindow() *OpenWindow {
	return &OpenWindow{
		item:           systray.AddMenuItem("", ""),
		chanOpenWindow: make(chan struct{}, 1),
	}
}

func (ow *OpenWindow) SetLocale(locale map[string]string) *OpenWindow {
	ow.item.SetTitle(locale["open_window"])
	ow.item.SetTooltip(locale["open_window"])
	return ow
}

func (ow *OpenWindow) Watch(listener OpenWindowListener) *OpenWindow {
	go func() {
		for {
			select {
			case <-ow.item.ClickedCh:
				listener.OnOpenWindow()
			case <-ow.chanOpenWindow:
				return
			}
		}
	}()
	return ow
}

func (ow *OpenWindow) StopWatch() *OpenWindow {
	close(ow.chanOpenWindow)
	return ow
}

func (ow *OpenWindow) Trigger() *OpenWindow {
	ow.item.ClickedCh <- struct{}{}
	return ow
}
