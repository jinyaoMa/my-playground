package backend

import (
	"context"
	"my-playground/backend/server"
	"my-playground/backend/tray"
	"my-playground/backend/utils"
)

// App struct
type App struct {
	ctx    context.Context
	config *Config
}

/* START - Public methods binded to wails frontend */

func (a *App) ChangeLanguage(lang string) {
	tray.My().ChangeLanguage(lang)
	utils.Logger(PkgName).Printf("WAILS/App ChangeLanguage(%s) \n", lang)
}

/* END - Public methods binded to wails frontend */

// startup is called at application startup
// startup 在应用程序启动时调用
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// 在这里执行初始化设置
	a.ctx = ctx
	a.config = MyConfig(ctx)
	server.My().SetConfig(a.config.Server)
	tray.My().SetConfig(a.config.Tray)

	tray.Start()
	utils.Logger(PkgName).Println("WAILS START UP")
}

// domReady is called after the front-end dom has been loaded
// domReady 在前端Dom加载完毕后调用
func (a *App) domReady(ctx context.Context) {
	// Add your action here
	// 在这里添加你的操作
	utils.Logger(PkgName).Println("WAILS DOM READY")
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
// beforeClose在单击窗口关闭按钮或调用runtime.Quit即将退出应用程序时被调用.
// 返回 true 将导致应用程序继续，false 将继续正常关闭。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	utils.Logger(PkgName).Println("WAILS BEFORE CLOSE")
	return false
}

// shutdown is called at application termination
// 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 在此处做一些资源释放的操作
	tray.Stop()
	utils.Logger(PkgName).Println("WAILS SHUTDOWN")
}

func (a *App) suspend() {
	// Add your action here
	// 在这里添加你的操作
	utils.Logger(PkgName).Println("WAILS SUSPEND")
}

func (a *App) resume() {
	// Add your action here
	// 在这里添加你的操作
	utils.Logger(PkgName).Println("WAILS RESUME")
}
