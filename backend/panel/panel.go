package panel

import (
	"embed"
	"log"
	"my-playground/backend/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func Run(frontend embed.FS) {
	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := app.Lication()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             "My Services",
		Width:             900,
		Height:            600,
		MinWidth:          900,
		MinHeight:         600,
		MaxWidth:          1200,
		MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: false,
		RGBA:              &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Assets:            frontend,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.Startup,
		OnDomReady:        app.DomReady,
		OnBeforeClose:     app.BeforeClose,
		OnShutdown:        app.Shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: true,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			OnSuspend: func() {
			},
			OnResume: func() {
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}