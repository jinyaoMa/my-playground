package tray

import (
	"context"
	"embed"
	"fmt"
	"my-playground/backend/model"
	"my-playground/backend/server"
	"my-playground/backend/tray/menus"
	"my-playground/backend/utils"
	"strings"
	"sync"

	"fyne.io/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icons/icon.ico
var icon []byte

//go:embed icons/open-window.ico
var iconOpenWindow []byte

//go:embed icons/api-start.ico
var iconApiStart []byte

//go:embed icons/api-stop.ico
var iconApiStop []byte

//go:embed locales/en.json
//go:embed locales/zh.json
var locales embed.FS

const (
	PkgName = "tray"

	CfgNameLanguage = "Tray.Language"
	CfgNameTheme    = "Tray.Theme"
)

var (
	tray *Tray
)

func init() {
	tray = &Tray{}
	tray.start, tray.stop = systray.RunWithExternalLoop(tray.onReady, tray.onQuit)
}

type Config struct {
	Context  context.Context // bind wails context
	Language string
	Theme    string
}

type Tray struct {
	config      *Config
	start, stop func()    // start/stop tray
	once        sync.Once // make sure only run start tray functions once
	locale      map[string]string
	openWindow  *menus.OpenWindow
	apiService  *menus.ApiService
	language    *menus.Language
	theme       *menus.Theme
	quit        *menus.Quit
}

func My() *Tray {
	return tray
}

func (t *Tray) SetConfig(cfg *Config) *Tray {
	t.config = cfg
	return t.ChangeLanguage(cfg.Language).
		ChangeTheme(cfg.Theme)
}

func (t *Tray) ChangeLanguage(lang string) *Tray {
	switch lang {
	default:
		t.language.ClickChinese()
	case "en":
		t.language.ClickEnglish()
	}
	return t
}

func (t *Tray) ChangeTheme(theme string) *Tray {
	switch theme {
	default:
		t.theme.ClickLight()
	case "dark":
		t.theme.ClickDark()
	}
	return t
}

func (t *Tray) Start() *Tray {
	t.once.Do(t.start)
	return t
}

func (t *Tray) Stop() *Tray {
	t.stop()
	return t
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
		SetIconStart(iconApiStart).
		SetIconStop(iconApiStop).
		Watch(menus.ApiServiceListener{
			OnStart: func() bool {
				return server.My().Start()
			},
			OnStop: func() bool {
				return server.My().Stop()
			},
			OnOpenSwagger: func() {
				runtime.BrowserOpenURL(
					t.config.Context,
					fmt.Sprintf("https://localhost%s/swagger/index.html", server.My().GetHttpsPort()),
				)
			},
		})

	systray.AddSeparator()

	t.language = menus.
		NewLanguage().
		Watch(menus.LanguageListener{
			OnLanguageChanged: func(filename string) bool {
				runtime.EventsEmit(t.config.Context, "onLanguageChanged", t.locale2Lang(filename))
				return t.updateLocales(filename)
			},
		})

	systray.AddSeparator()

	t.theme = menus.
		NewTheme().
		Watch(menus.ThemeListener{
			OnThemeChanged: func(theme string) bool {
				runtime.EventsEmit(t.config.Context, "OnThemeChanged", theme)
				return t.updateTheme(theme)
			},
		})

	systray.AddSeparator()

	t.quit = menus.
		NewQuit().
		Watch(menus.QuitListener{
			OnQuit: func() {
				dialog, err := runtime.MessageDialog(t.config.Context, runtime.MessageDialogOptions{
					Type:          runtime.QuestionDialog,
					Title:         t.locale["quitDialog.title"],
					Message:       t.locale["quitDialog.message"],
					Buttons:       []string{},
					DefaultButton: t.locale["quitDialog.defaultButtun"],
					CancelButton:  t.locale["quitDialog.cancelButton"],
				})
				if err != nil {
					utils.Logger(PkgName).Fatalf("fail to open quit dialog: %+v\n", err)
				}
				if dialog == "Yes" { // when default button => "Yes" is clicked
					t.Stop()
				}
			},
		})
}

func (t *Tray) onQuit() {
	server.My().Stop()

	{ // end menus properly
		t.openWindow.StopWatch()
		t.apiService.StopWatch()
		t.language.StopWatch()
		t.theme.StopWatch()
		t.quit.StopWatch()
	}

	runtime.Quit(t.config.Context)
}

func (t *Tray) updateTheme(theme string) bool {
	switch theme {
	default:
		return false
	case "light":
		runtime.WindowSetLightTheme(t.config.Context)
	case "dark":
		runtime.WindowSetDarkTheme(t.config.Context)
	}

	t.config.Theme = theme
	option := model.MpOption{
		Name: CfgNameTheme,
	}
	result := option.Update(t.config.Theme)
	if result.Error != nil {
		utils.Logger(PkgName).Fatalf("failed to update theme option: %+v\n", result.Error)
	}

	return true
}

func (t *Tray) updateLocales(filename string) bool {
	rawJson, _ := locales.ReadFile(filename)
	t.locale = utils.GetLocaleFromJSON(rawJson)

	runtime.WindowSetTitle(t.config.Context, t.locale["appname"])
	systray.SetTitle(t.locale["appname"])
	systray.SetTooltip(t.locale["appname"])
	t.openWindow.SetLocale(t.locale)
	t.apiService.SetLocale(t.locale)
	t.language.SetLocale(t.locale)
	t.theme.SetLocale(t.locale)
	t.quit.SetLocale(t.locale)

	t.config.Language = t.locale2Lang(filename)
	option := model.MpOption{
		Name: CfgNameLanguage,
	}
	result := option.Update(t.config.Language)
	if result.Error != nil {
		utils.Logger(PkgName).Fatalf("failed to update language option: %+v\n", result.Error)
	}

	return true
}

func (t *Tray) locale2Lang(filename string) string {
	return strings.Replace(strings.Replace(filename, "locales/", "", 1), ".json", "", 1)
}
