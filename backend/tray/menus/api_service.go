package menus

import (
	"github.com/getlantern/systray"
)

type ApiServiceListener struct {
	OnStart func() (ok bool)
	OnStop  func() (ok bool)
}

type ApiService struct {
	isEnabled bool
	start     *systray.MenuItem
	stop      *systray.MenuItem
	chanStop  chan struct{}
}

func NewApiService() *ApiService {
	return &ApiService{
		start:    systray.AddMenuItem("", ""),
		stop:     systray.AddMenuItem("", ""),
		chanStop: make(chan struct{}, 1),
	}
}

func (as *ApiService) SetIconStart(icon []byte) *ApiService {
	as.start.SetIcon(icon)
	return as
}

func (as *ApiService) SetIconStop(icon []byte) *ApiService {
	as.stop.SetIcon(icon)
	return as
}

func (as *ApiService) SetLocale(locale map[string]string) *ApiService {
	as.start.SetTitle(locale["apiService.start"])
	as.start.SetTooltip(locale["apiService.start"])
	as.stop.SetTitle(locale["apiService.stop"])
	as.stop.SetTooltip(locale["apiService.stop"])
	if as.isEnabled {
		as.start.Hide()
		as.stop.Show()
	} else {
		as.start.Show()
		as.stop.Hide()
	}
	return as
}

func (as *ApiService) Watch(listener ApiServiceListener) *ApiService {
	go func() {
		for {
			select {
			case <-as.start.ClickedCh:
				if listener.OnStart() {
					as.start.Hide()
					as.stop.Show()
					as.isEnabled = true
				}
			case <-as.stop.ClickedCh:
				if listener.OnStop() {
					as.start.Show()
					as.stop.Hide()
					as.isEnabled = false
				}
			case <-as.chanStop:
				return
			}
		}
	}()
	return as
}

func (as *ApiService) StopWatch() *ApiService {
	close(as.chanStop)
	return as
}

func (as *ApiService) ClickStart() *ApiService {
	as.start.ClickedCh <- struct{}{}
	return as
}

func (as *ApiService) ClickStop() *ApiService {
	as.stop.ClickedCh <- struct{}{}
	return as
}
