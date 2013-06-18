package glfw3

//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
//#include "callback.h"
import "C"
import (
	"unsafe"
)

func (w *Window) SetPosCallback(callback IntPosCallback) {
	w.posCallback = &callback
	C.cSetPosCallback(w.pointer)
}
func (w *Window) SetSizeCallback(callback IntPosCallback) {
	w.sizeCallback = &callback
	C.cSetSizeCallback(w.pointer)
}
func (w *Window) SetCloseCallback(callback EventCallback) {
	w.closeCallback = &callback
	C.cSetCloseCallback(w.pointer)
}
func (w *Window) SetRefreshCallback(callback EventCallback) {
	w.refreshCallback = &callback
	C.cSetRefreshCallback(w.pointer)
}
func (w *Window) SetFocusCallback(callback StatusCallback) {
	w.focusCallback = &callback
	C.cSetFocusCallback(w.pointer)
}
func (w *Window) SetIconifyCallback(callback StatusCallback) {
	w.iconifyCallback = &callback
	C.cSetIconifyCallback(w.pointer)
}
/*
func (w *Window) SetFramebufferSizeCallback(callback SizeCallback) {
	w.framebufferSizeCallback = &callback
	C.cSetFramebufferSizeCallback(w.pointer)
}
*/


//export goPosCallback
func goPosCallback(window unsafe.Pointer, posX, posY C.int) {
	callback := getGoWindow(unsafe.Pointer(window)).posCallback
	if (callback != nil) {
		(*callback)(int(posX), int(posY))
	}
}
//export goSizeCallback
func goSizeCallback(window unsafe.Pointer, width, height C.int) {
	callback := getGoWindow(unsafe.Pointer(window)).sizeCallback
	if (callback != nil) {
		(*callback)(int(width), int(height))
	}
}
//export goCloseCallback
func goCloseCallback(window unsafe.Pointer) {
	callback := getGoWindow(unsafe.Pointer(window)).closeCallback
	if (callback != nil) {
		(*callback)()
	}
}
//export goRefreshCallback
func goRefreshCallback(window unsafe.Pointer) {
	callback := getGoWindow(unsafe.Pointer(window)).refreshCallback
	if (callback != nil) {
		(*callback)()
	}
}
//export goFocusCallback
func goFocusCallback(window unsafe.Pointer, focused C.int) {
	callback := getGoWindow(unsafe.Pointer(window)).focusCallback
	if (callback != nil) {
		(*callback)(int(focused) != 0)
	}
}
//export goIconifyCallback
func goIconifyCallback(window unsafe.Pointer, iconified C.int) {
	callback := getGoWindow(unsafe.Pointer(window)).iconifyCallback
	if (callback != nil) {
		(*callback)(int(iconified) != 0)
	}
}
/*
//export goFramebufferSizeCallback
func goFramebufferSizeCallback(window unsafe.Pointer, width, height C.int) {
	callback := getGoWindow(unsafe.Pointer(window)).framebufferSizeCallback
	if (callback != nil) {
		(*callback)(int(width), int(height))
	}
}
*/
