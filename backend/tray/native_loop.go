package tray

import (
	"fmt"
	"unsafe"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows"
)

// Code from github.com/getlantern/systray
// Modified to test abnormal ending of nativeLoop

var (
	u32               = windows.NewLazySystemDLL("User32.dll")
	pDispatchMessage  = u32.NewProc("DispatchMessageW")
	pGetMessage       = u32.NewProc("GetMessageW")
	pTranslateMessage = u32.NewProc("TranslateMessage")
)

func (t *Tray) customNativeLoop() {
	// Main message pump.
	m := &struct {
		WindowHandle windows.Handle
		Message      uint32
		Wparam       uintptr
		Lparam       uintptr
		Time         uint32
		// The POINT structure defines the x- and y- coordinates of a point.
		// https://msdn.microsoft.com/en-us/library/windows/desktop/dd162805(v=vs.85).aspx
		Pt struct {
			X, Y int32
		}
	}{}
	for {
		ret, _, err := pGetMessage.Call(uintptr(unsafe.Pointer(m)), 0, 0, 0)

		// If the function retrieves a message other than WM_QUIT, the return value is nonzero.
		// If the function retrieves the WM_QUIT message, the return value is zero.
		// If there is an error, the return value is -1
		// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644936(v=vs.85).aspx
		switch int32(ret) {
		case -1:
			runtime.MessageDialog(t.ctx, runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Title:   "Error",
				Message: fmt.Sprintf("Error at message loop: %v", err),
			})
			return
		case 0:
			runtime.MessageDialog(t.ctx, runtime.MessageDialogOptions{
				Type:    runtime.WarningDialog,
				Title:   "Warning",
				Message: "SYSTRAY NATIVELOOP (0)",
			})
			return
		default:
			pTranslateMessage.Call(uintptr(unsafe.Pointer(m)))
			pDispatchMessage.Call(uintptr(unsafe.Pointer(m)))
		}
	}
}
