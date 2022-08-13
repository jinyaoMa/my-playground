package menus

import (
	"github.com/getlantern/systray"
)

type LanguageListener struct {
	OnLanguageChanged func(filename string) (ok bool)
}

type Language struct {
	item     *systray.MenuItem
	chinese  *systray.MenuItem
	english  *systray.MenuItem
	chanStop chan struct{}
}

func NewLanguage() *Language {
	return &Language{
		item:     systray.AddMenuItem("", ""),
		chinese:  systray.AddMenuItem("", ""),
		english:  systray.AddMenuItem("", ""),
		chanStop: make(chan struct{}, 1),
	}
}

func (l *Language) SetIcon(icon []byte) *Language {
	l.item.SetIcon(icon)
	return l
}

func (l *Language) SetIconChinese(icon []byte) *Language {
	l.chinese.SetIcon(icon)
	return l
}

func (l *Language) SetIconEnglish(icon []byte) *Language {
	l.english.SetIcon(icon)
	return l
}

func (l *Language) SetLocale(locale map[string]string) *Language {
	l.item.SetTitle(locale["displayLanguage"])
	l.item.SetTooltip(locale["displayLanguage"])
	l.item.Disable()
	l.chinese.SetTitle(locale["chinese"])
	l.chinese.SetTooltip(locale["chinese"])
	l.english.SetTitle(locale["english"])
	l.english.SetTooltip(locale["english"])
	return l
}

func (l *Language) Watch(listener LanguageListener) *Language {
	go func() {
		for {
			select {
			case <-l.chinese.ClickedCh:
				if listener.OnLanguageChanged("locales/zh.json") {
					l.chinese.Check()
					l.english.Uncheck()
				}
			case <-l.english.ClickedCh:
				if listener.OnLanguageChanged("locales/en.json") {
					l.chinese.Uncheck()
					l.english.Check()
				}
			case <-l.chanStop:
				return
			}
		}
	}()
	return l
}

func (l *Language) StopWatch() *Language {
	close(l.chanStop)
	return l
}

func (l *Language) ClickChinese() *Language {
	l.chinese.ClickedCh <- struct{}{}
	return l
}

func (l *Language) ClickEnglish() *Language {
	l.english.ClickedCh <- struct{}{}
	return l
}
