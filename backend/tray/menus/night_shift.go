package menus

import (
	"fyne.io/systray"
)

type NightShiftListener struct {
	OnNightShift func(isNightShift bool) (ok bool)
}

type NightShift struct {
	isNightShift bool
	item         *systray.MenuItem
	chanStop     chan struct{}
}

func NewNightShift() *NightShift {
	return &NightShift{
		item:     systray.AddMenuItem("", ""),
		chanStop: make(chan struct{}, 1),
	}
}

func (ns *NightShift) SetIcon(icon []byte) *NightShift {
	ns.item.SetIcon(icon)
	return ns
}

func (ns *NightShift) SetLocale(locale map[string]string) *NightShift {
	ns.item.SetTitle(locale["nightShift"])
	ns.item.SetTooltip(locale["nightShift"])
	if ns.isNightShift {
		ns.item.Check()
	} else {
		ns.item.Uncheck()
	}
	return ns
}

func (ns *NightShift) Watch(listener NightShiftListener) *NightShift {
	go func() {
		for {
			select {
			case <-ns.item.ClickedCh:
				isNightShift := !ns.isNightShift
				if listener.OnNightShift(isNightShift) {
					if isNightShift {
						ns.item.Check()
					} else {
						ns.item.Uncheck()
					}
					ns.isNightShift = isNightShift
				}
			case <-ns.chanStop:
				return
			}
		}
	}()
	return ns
}

func (ns *NightShift) StopWatch() *NightShift {
	close(ns.chanStop)
	return ns
}

func (ns *NightShift) Click() *NightShift {
	ns.item.ClickedCh <- struct{}{}
	return ns
}
