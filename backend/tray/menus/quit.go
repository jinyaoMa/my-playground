package menus

import (
	"github.com/getlantern/systray"
)

type QuitListener struct {
	OnQuit func()
}

type QuitLocale struct {
	Title   string
	Tooltip string
}

type Quit struct {
	menu     *systray.MenuItem
	chanQuit chan struct{}
}

func (q *Quit) Init(locale QuitLocale) {
	q.menu = systray.AddMenuItem(locale.Title, locale.Tooltip)
	q.chanQuit = make(chan struct{}, 1)
}

func (q *Quit) UpdateLocale(locale QuitLocale) {
	q.menu.SetTitle(locale.Title)
	q.menu.SetTooltip(locale.Tooltip)
}

func (q *Quit) Watch(listener QuitListener) {
	go func() {
		for {
			select {
			case <-q.menu.ClickedCh:
				listener.OnQuit()
			case <-q.chanQuit:
				return
			}
		}
	}()
}

func (q *Quit) StopWatch() {
	q.chanQuit <- struct{}{}
}
