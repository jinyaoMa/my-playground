package tray

import (
	"context"
	"embed"
	"my-playground/backend/tray/menus"
	"my-playground/backend/utils"
	"strings"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icons/icon.ico
var favicon []byte

//go:embed locales/en.json
//go:embed locales/zh.json
var locales embed.FS

var tray *Tray

type Tray struct {
	ctx        context.Context
	openWindow *menus.OpenWindow
	language   *menus.Language
	quit       *menus.Quit
}

func Setup(ctx context.Context) {
	tray = &Tray{
		ctx: ctx,
	}
	go systray.Run(tray.onReady, tray.onExit)
}

func ChangeLanguage(filename string) {
	switch {
	default:
		tray.language.ClickChinese()
	case strings.Contains(filename, "en"):
		tray.language.ClickEnglish()
	}
}

func (t *Tray) onReady() {
	systray.SetIcon(favicon)

	t.openWindow = menus.
		NewOpenWindow().
		SetIcon(favicon).
		Watch(menus.OpenWindowListener{
			OnOpenWindow: func() {
				runtime.WindowShow(t.ctx)
			},
		})

	systray.AddSeparator()

	t.language = menus.
		NewLanguage().
		Watch(menus.LanguageListener{
			OnLanguageChanged: func(filename string) bool {
				runtime.EventsEmit(t.ctx, "onLanguageChanged", filename)
				t.updateLocales(filename)
				return true
			},
		})

	systray.AddSeparator()

	t.quit = menus.
		NewQuit().
		Watch(menus.QuitListener{
			OnQuit: func() {
				dialog, _ := runtime.MessageDialog(t.ctx, runtime.MessageDialogOptions{
					Type:          runtime.QuestionDialog,
					Title:         "Quit?",
					Message:       "Are you sure you want to quit?",
					Buttons:       []string{},
					DefaultButton: "Yes",
					CancelButton:  "No",
				})
				if dialog == "Yes" {
					systray.Quit()
				}
			},
		})

	t.language.ClickChinese()
}

func (t *Tray) onExit() {
	t.openWindow.StopWatch()
	t.quit.StopWatch()
	runtime.Quit(t.ctx)
}

func (t *Tray) updateLocales(filename string) {
	rawJson, _ := locales.ReadFile(filename)
	locale := utils.GetLocaleFromJSON(rawJson)

	systray.SetTitle(locale["appname"])
	systray.SetTooltip(locale["appname"])
	t.openWindow.SetLocale(locale)
	t.language.SetLocale(locale)
	t.quit.SetLocale(locale)
}
