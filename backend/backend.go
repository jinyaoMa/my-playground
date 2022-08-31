package backend

import (
	"embed"
	"my-playground/backend/utils"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed .frontend
var frontend embed.FS

const (
	PkgName = "backend"
)

func RunApp() (app *App) {
	app = &App{}

	var wailsLogger logger.Logger
	if utils.IsDev() {
		wailsLogger = logger.NewDefaultLogger()
	} else {
		wailsLogger = logger.NewFileLogger(utils.GetExecutablePath("wails.log"))
	}

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:              "My Playground",
		Width:              800,
		Height:             600,
		DisableResize:      false,
		Fullscreen:         false,
		Frameless:          true,
		MinWidth:           800,
		MinHeight:          600,
		MaxWidth:           -1,
		MaxHeight:          -1,
		StartHidden:        false,
		HideWindowOnClose:  true,
		AlwaysOnTop:        false,
		BackgroundColour:   &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Assets:             frontend,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             wailsLogger,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          app.startup,
		OnDomReady:         app.domReady,
		OnShutdown:         app.shutdown,
		OnBeforeClose:      app.beforeClose,
		Bind:               []interface{}{app},
		WindowStartState:   options.Normal,
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 true,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			OnSuspend:                         app.suspend,
			OnResume:                          app.resume,
		},
	})

	if err != nil {
		utils.Logger(PkgName).Fatalf("fail to run wails: %+v\n", err)
	}
	return
}
