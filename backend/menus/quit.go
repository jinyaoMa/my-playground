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
	item     *systray.MenuItem
	chanQuit chan struct{}
}

func NewQuit(locale QuitLocale) *Quit {
	return &Quit{
		item:     systray.AddMenuItem(locale.Title, locale.Tooltip),
		chanQuit: make(chan struct{}, 1),
	}
}

func (q *Quit) UpdateLocale(locale QuitLocale) *Quit {
	q.item.SetTitle(locale.Title)
	q.item.SetTooltip(locale.Tooltip)
	return q
}

func (q *Quit) Watch(listener QuitListener) *Quit {
	go func() {
		for {
			select {
			case <-q.item.ClickedCh:
				listener.OnQuit()
			case <-q.chanQuit:
				return
			}
		}
	}()
	return q
}

func (q *Quit) StopWatch() *Quit {
	close(q.chanQuit)
	return q
}

func (q *Quit) Trigger() *Quit {
	q.item.ClickedCh <- struct{}{}
	return q
}
