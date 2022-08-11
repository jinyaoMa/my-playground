package backend

import (
	"context"
	"embed"
	"log"
	"my-playground/backend/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed .assets
var assets embed.FS

type Backend struct {
	ctx context.Context
}

func Run() {
	b := &Backend{}
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
		Assets:             assets,
		AssetsHandler:      nil,
		Menu:               nil,
		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		OnStartup:          b.startup,
		OnDomReady:         b.domReady,
		OnShutdown:         b.shutdown,
		OnBeforeClose:      b.beforeClose,
		Bind:               []interface{}{b},
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
			OnSuspend:                         b.suspend,
			OnResume:                          b.resume,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

// startup is called at application startup
// startup 在应用程序启动时调用
func (b *Backend) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	b.ctx = ctx
	app.Lication(ctx)
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (b *Backend) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (b *Backend) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (b *Backend) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
	app.Lication().QuitTray()
}

func (b *Backend) suspend() {
	// Add your action here
	// 在这里添加你的操作
}

func (b *Backend) resume() {
	// Add your action here
	// 在这里添加你的操作
}

/* Public methods binded to wails frontend */
