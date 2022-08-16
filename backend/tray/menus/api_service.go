package menus

import (
	"fyne.io/systray"
)

type ApiServiceListener struct {
	OnStart       func() (ok bool)
	OnStop        func() (ok bool)
	OnOpenSwagger func()
}

type ApiService struct {
	isEnabled bool
	start     *systray.MenuItem
	stop      *systray.MenuItem
	swagger   *systray.MenuItem
	chanStop  chan struct{}
}

func NewApiService() *ApiService {
	return &ApiService{
		start:    systray.AddMenuItem("", ""),
		stop:     systray.AddMenuItem("", ""),
		swagger:  systray.AddMenuItem("", ""),
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

func (as *ApiService) SetIconSwagger(icon []byte) *ApiService {
	as.swagger.SetIcon(icon)
	return as
}

func (as *ApiService) SetLocale(locale map[string]string) *ApiService {
	as.start.SetTitle(locale["apiService.start"])
	as.start.SetTooltip(locale["apiService.start"])
	as.stop.SetTitle(locale["apiService.stop"])
	as.stop.SetTooltip(locale["apiService.stop"])
	as.swagger.SetTitle(locale["apiService.swagger"])
	as.swagger.SetTooltip(locale["apiService.swagger"])
	if as.isEnabled {
		as.start.Hide()
		as.stop.Show()
		as.swagger.Show()
	} else {
		as.start.Show()
		as.stop.Hide()
		as.swagger.Hide()
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
					as.swagger.Show()
					as.isEnabled = true
				}
			case <-as.stop.ClickedCh:
				if listener.OnStop() {
					as.start.Show()
					as.stop.Hide()
					as.swagger.Hide()
					as.isEnabled = false
				}
			case <-as.swagger.ClickedCh:
				listener.OnOpenSwagger()
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

func (as *ApiService) ClickOpenSwagger() *ApiService {
	as.swagger.ClickedCh <- struct{}{}
	return as
}
