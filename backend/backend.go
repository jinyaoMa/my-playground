package backend

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed .assets
var assets embed.FS

func Run() {
	app := &App{}

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
		Assets:             assets,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             nil,
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
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			TranslucencyType:                  windows.Auto,
			OnSuspend:                         app.suspend,
			OnResume:                          app.resume,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
