package menus

import (
	"fyne.io/systray"
)

type QuitListener struct {
	OnQuit func()
}

type Quit struct {
	item     *systray.MenuItem
	chanStop chan struct{}
}

func NewQuit() *Quit {
	return &Quit{
		item:     systray.AddMenuItem("", ""),
		chanStop: make(chan struct{}, 1),
	}
}

func (q *Quit) SetIcon(icon []byte) *Quit {
	q.item.SetIcon(icon)
	return q
}

func (q *Quit) SetLocale(locale map[string]string) *Quit {
	q.item.SetTitle(locale["quit"])
	q.item.SetTooltip(locale["quit"])
	return q
}

func (q *Quit) Watch(listener QuitListener) *Quit {
	go func() {
		for {
			select {
			case <-q.item.ClickedCh:
				listener.OnQuit()
			case <-q.chanStop:
				return
			}
		}
	}()
	return q
}

func (q *Quit) StopWatch() *Quit {
	close(q.chanStop)
	return q
}

func (q *Quit) Click() *Quit {
	q.item.ClickedCh <- struct{}{}
	return q
}
