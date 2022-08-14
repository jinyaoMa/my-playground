package tray

import (
	"context"
	"embed"
	"my-playground/backend/server"
	"my-playground/backend/tray/menus"
	"my-playground/backend/utils"
	"strings"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icons/tray.icon.ico
var icon []byte

//go:embed icons/open-window.icon.ico
var iconOpenWindow []byte

//go:embed icons/api-start.icon.ico
var iconApiStart []byte

//go:embed icons/api-stop.icon.ico
var iconApiStop []byte

//go:embed locales/en.json
//go:embed locales/zh.json
var locales embed.FS

var (
	tray   *Tray
	config *Config
)

type Tray struct {
	ctx        context.Context // bind wails context
	locale     map[string]string
	openWindow *menus.OpenWindow
	apiService *menus.ApiService
	language   *menus.Language
	quit       *menus.Quit
}

type Config struct {
	WailsCtx context.Context
	Language string
}

func Setup(cfg *Config) {
	config = cfg

	tray = &Tray{
		ctx: config.WailsCtx,
	}
	go systray.Run(tray.onReady, tray.onQuit)
}

func ChangeLanguage(lang string) {
	switch lang {
	default:
		tray.language.ClickChinese()
	case "en":
		tray.language.ClickEnglish()
	}
}

func (t *Tray) onReady() {
	systray.SetIcon(icon)

	t.openWindow = menus.
		NewOpenWindow().
		SetIcon(iconOpenWindow).
		Watch(menus.OpenWindowListener{
			OnOpenWindow: func() {
				runtime.WindowShow(t.ctx)
			},
		})

	systray.AddSeparator()

	t.apiService = menus.
		NewApiService().
		SetIconStart(iconApiStop).
		SetIconStop(iconApiStart).
		Watch(menus.ApiServiceListener{
			OnStart: func() bool {
				return server.StartServer()
			},
			OnStop: func() bool {
				return server.StopServer()
			},
		})

	systray.AddSeparator()

	t.language = menus.
		NewLanguage().
		Watch(menus.LanguageListener{
			OnLanguageChanged: func(filename string) bool {
				runtime.EventsEmit(t.ctx, "onLanguageChanged", t.locale2Lang(filename))
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
					Title:         t.locale["quitDialog.title"],
					Message:       t.locale["quitDialog.message"],
					Buttons:       []string{},
					DefaultButton: t.locale["quitDialog.defaultButtun"],
					CancelButton:  t.locale["quitDialog.cancelButton"],
				})
				if dialog == "Yes" { // when default button => "Yes" is clicked
					systray.Quit()
				}
			},
		})

	ChangeLanguage(config.Language)
	// t.apiService.ClickStart()
}

func (t *Tray) onQuit() {
	t.openWindow.StopWatch()
	t.quit.StopWatch()
	server.StopServer()
	runtime.Quit(t.ctx)
}

func (t *Tray) updateLocales(filename string) {
	rawJson, _ := locales.ReadFile(filename)
	t.locale = utils.GetLocaleFromJSON(rawJson)

	runtime.WindowSetTitle(t.ctx, t.locale["appname"])
	systray.SetTitle(t.locale["appname"])
	systray.SetTooltip(t.locale["appname"])
	t.openWindow.SetLocale(t.locale)
	t.apiService.SetLocale(t.locale)
	t.language.SetLocale(t.locale)
	t.quit.SetLocale(t.locale)
}

func (t *Tray) locale2Lang(filename string) string {
	return strings.Replace(strings.Replace(filename, "locales/", "", 1), ".json", "", 1)
}
