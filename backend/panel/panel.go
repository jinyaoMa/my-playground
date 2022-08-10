package panel

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func Run(frontend embed.FS) {
	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:              "",
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
		HideWindowOnClose:  false,
		AlwaysOnTop:        false,
		BackgroundColour:   &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Assets:             frontend,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          app.startup,
		OnDomReady:         app.domReady,
		OnShutdown:         app.shutdown,
		OnBeforeClose:      app.beforeClose,
		Bind: []interface{}{
			app,
		},
		WindowStartState: options.Normal,
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			OnSuspend: func() {
			},
			OnResume: func() {
			}},
	})

	if err != nil {
		log.Fatal(err)
	}
}
