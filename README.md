# my-playground

A collection of services.

## Ports Used in Development

* `10000` for Wails backend, edit `wails.json` to configure
* `10001` for Wails frontend, edit `frontend/app/vite.config.js` to configure
* `10002` for documentation, edit `frontend/docs/.vitepress/config.js` to configure
* `10080` and `10443` for in-process web server, edit `/backend/config.go` to configure

## Environment, Tools and Dependencies

> I am using Windows 10 Home 21H1 19043.1826 with [VS Code + GCC](https://code.visualstudio.com/docs/cpp/config-mingw).

| Name              | Version               | Link                                                           |
| :---------------- | :-------------------- | :------------------------------------------------------------- |
| Git               | 2.37+                 | https://git-scm.com/                                           |
| GO                | 1.18+                 | https://go.dev/dl/                                             |
| Node (with NPM)   | 16.16+ (NPM v8.15.1+) | https://nodejs.org/                                            |
| WebView2          | 103.0.1264.62+        | https://developer.microsoft.com/en-us/microsoft-edge/webview2/ |
| UPX               | 3.96+                 | https://upx.github.io/                                         |
| Wails             | v2.0.0-beta.40+       | https://wails.io/                                              |
| Lerna             | 5.3+                  | https://lerna.js.org/                                          |
| Vite              | 2.9+                  | https://vitejs.dev/                                            |
| VitePress         | 1.0.0-alpha.4+        | https://vitepress.vuejs.org/                                   |
| Vue               | 3.2+                  | https://vuejs.org/                                             |
| Systray (fyne.io) | 1.10.1+               | https://github.com/fyne-io/systray                             |
| Gorm              | 1.23.8+               | https://gorm.io/                                               |
| SQLite Driver     | 1.3.6+                | https://gorm.io/                                               |
| Gin               | 1.8.1+                | https://gin-gonic.com/                                         |
| Swaggo            | 1.8.4+                | https://github.com/swaggo/swag                                 |
| Gin Swagger       | 1.5.2+                | https://github.com/swaggo/gin-swagger                          |
| Air               | 1.40.4+               | https://github.com/cosmtrek/air                                |

## Backend Code Script Order Norms

``` go
//go:embed or other variables with special comments
var PubFs embed.FS // public
//go:embed or other variables with special comments
var PriFs embed.FS // private

const (
  // Constant variable (must be public)
  // Variable name starts with 3-letter prefix
  // Capital first letter of each word
  CstIntValue = iota
  CstIntValue1
  // ...
)

var (
  // Public variable
  // Capital first letter of each word
  PubSimpleFunc func()

  // Private variable
  // Capital first letter of each word except the first one
  priIntValue int
)

func init() {}
func main() {}

// Util/Middleware/Handler functions (must be public)
func UtilFunction() {}

// Interface starts with Capital 'I'
// Capital first letter of each word
type ISimpleInterface interface {}

// Struct
// Capital first letter of each word
type SimpleStruct struct {}

// Constructor function
func Constructor() {}

// Public struct functions
// SimpleStruct => ss
func (ss *SimpleStruct) PublicFunction() {}

// Private struct functions
func (ss *SimpleStruct) privateFunction() {}

// Other structs...
type AnotherStruct struct {}
func Constructor() {}
// AnotherStruct => as
func (as *AnotherStruct) PublicFunction() {}
func (as *AnotherStruct) privateFunction() {}
// ...
```

``` go
// Log Norms
// Printf/Println: Capital first letter of each word OR capitalize all characters
Printf("Wails/App ChangeLanguage(%s)\n", "zh") // remember to add newline '\n' at the end
Println("WAILS SHUTDOWN")
// Fatalf: all characters lowercase, end with ': %+v\n' and error
Fatalf("fail to do something: %+v\n", err)
```
