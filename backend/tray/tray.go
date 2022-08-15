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
	tray        *Tray
	Start, Stop func() // start/stop tray
)

const (
	CfgNameLanguage = "Tray.Language"
)

type Config struct {
	Context  context.Context // bind wails context
	Language string
}

type Tray struct {
	config     *Config
	locale     map[string]string
	openWindow *menus.OpenWindow
	apiService *menus.ApiService
	language   *menus.Language
	quit       *menus.Quit
}

func init() {
	tray = &Tray{}
	Start, Stop = systray.RunWithExternalLoop(tray.onReady, tray.onQuit)
}

func SetConfig(cfg *Config) {
	tray.config = cfg
	ChangeLanguage(cfg.Language)
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
				runtime.WindowShow(t.config.Context)
			},
		})

	systray.AddSeparator()

	t.apiService = menus.
		NewApiService().
		SetIconStart(iconApiStop).
		SetIconStop(iconApiStart).
		Watch(menus.ApiServiceListener{
			OnStart: func() bool {
				return server.Start()
			},
			OnStop: func() bool {
				return server.Stop()
			},
		})

	systray.AddSeparator()

	t.language = menus.
		NewLanguage().
		Watch(menus.LanguageListener{
			OnLanguageChanged: func(filename string) bool {
				runtime.EventsEmit(t.config.Context, "onLanguageChanged", t.locale2Lang(filename))
				t.updateLocales(filename)
				return true
			},
		})

	systray.AddSeparator()

	t.quit = menus.
		NewQuit().
		Watch(menus.QuitListener{
			OnQuit: func() {
				dialog, _ := runtime.MessageDialog(t.config.Context, runtime.MessageDialogOptions{
					Type:          runtime.QuestionDialog,
					Title:         t.locale["quitDialog.title"],
					Message:       t.locale["quitDialog.message"],
					Buttons:       []string{},
					DefaultButton: t.locale["quitDialog.defaultButtun"],
					CancelButton:  t.locale["quitDialog.cancelButton"],
				})
				if dialog == "Yes" { // when default button => "Yes" is clicked
					runtime.Quit(t.config.Context)
				}
			},
		})
}

func (t *Tray) onQuit() {
	server.Stop()

	{ // end menus properly
		t.openWindow.StopWatch()
		t.apiService.StopWatch()
		t.language.StopWatch()
		t.quit.StopWatch()
	}
}

func (t *Tray) updateLocales(filename string) {
	rawJson, _ := locales.ReadFile(filename)
	t.locale = utils.GetLocaleFromJSON(rawJson)

	runtime.WindowSetTitle(t.config.Context, t.locale["appname"])
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
