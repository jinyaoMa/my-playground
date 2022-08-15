package tray

import (
	"context"
	"embed"
	"log"
	"my-playground/backend/model"
	"my-playground/backend/server"
	"my-playground/backend/tray/menus"
	"my-playground/backend/utils"
	"strings"

	"fyne.io/systray"
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

const (
	CfgNameLanguage = "Tray.Language"
)

type Config struct {
	WailsCtx context.Context
	Language string
}

func Setup() (start, end func()) {
	tray = &Tray{}
	return systray.RunWithExternalLoop(tray.onReady, tray.onQuit)
}

func SetConfig(cfg *Config) {
	config = cfg
	tray.ctx = config.WailsCtx
}

func ChangeLanguage(lang string) {
	if tray == nil {
		return
	}

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
					runtime.Quit(t.ctx)
				}
			},
		})
}

func (t *Tray) onQuit() {
	server.StopServer()

	// end menus properly
	t.openWindow.StopWatch()
	t.apiService.StopWatch()
	t.language.StopWatch()
	t.quit.StopWatch()
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

	option := model.MpOption{
		Name: CfgNameLanguage,
	}
	result := option.Update(t.locale2Lang(filename))
	if result.Error != nil {
		log.Fatalln("failed to update language option")
	}
}

func (t *Tray) locale2Lang(filename string) string {
	return strings.Replace(strings.Replace(filename, "locales/", "", 1), ".json", "", 1)
}
