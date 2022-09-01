package menus

import (
	"fyne.io/systray"
)

type ThemeListener struct {
	OnThemeChanged func(theme string) (ok bool)
}

type Theme struct {
	item     *systray.MenuItem
	light    *systray.MenuItem
	dark     *systray.MenuItem
	chanStop chan struct{}
}

func NewTheme() *Theme {
	return &Theme{
		item:     systray.AddMenuItem("", ""),
		light:    systray.AddMenuItem("", ""),
		dark:     systray.AddMenuItem("", ""),
		chanStop: make(chan struct{}, 1),
	}
}

func (t *Theme) SetIcon(icon []byte) *Theme {
	t.item.SetIcon(icon)
	return t
}

func (t *Theme) SetIconLight(icon []byte) *Theme {
	t.light.SetIcon(icon)
	return t
}

func (t *Theme) SetIconDark(icon []byte) *Theme {
	t.dark.SetIcon(icon)
	return t
}

func (t *Theme) SetLocale(locale map[string]string) *Theme {
	t.item.SetTitle(locale["theme"])
	t.item.SetTooltip(locale["theme"])
	t.item.Disable()
	t.light.SetTitle(locale["theme.light"])
	t.light.SetTooltip(locale["theme.light"])
	t.dark.SetTitle(locale["theme.dark"])
	t.dark.SetTooltip(locale["theme.dark"])
	return t
}

func (t *Theme) Watch(listener ThemeListener) *Theme {
	go func() {
		for {
			select {
			case <-t.light.ClickedCh:
				if listener.OnThemeChanged("light") {
					t.light.Check()
					t.dark.Uncheck()
				}
			case <-t.dark.ClickedCh:
				if listener.OnThemeChanged("dark") {
					t.light.Uncheck()
					t.dark.Check()
				}
			case <-t.chanStop:
				return
			}
		}
	}()
	return t
}

func (t *Theme) StopWatch() *Theme {
	close(t.chanStop)
	return t
}

func (t *Theme) ClickLight() *Theme {
	t.light.ClickedCh <- struct{}{}
	return t
}

func (t *Theme) ClickDark() *Theme {
	t.dark.ClickedCh <- struct{}{}
	return t
}
