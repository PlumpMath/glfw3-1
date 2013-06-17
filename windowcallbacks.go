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



//SetSizeCallback
//SetCloseCallback
//SetRefreshCallback
//SetFocusCallback
//SetIconifyCallback
//SetFramebufferSizeCallback



type PosCallback func (int, int)
func (w *Window) SetPosCallback(callback PosCallback) {
	w.posCallback = &callback
	C.cSetPosCallback(w.pointer)
}
//export goPosCallback
func goPosCallback(window unsafe.Pointer, posX, posY C.int) {
	callback := getGoWindow(unsafe.Pointer(window)).posCallback
	if (callback != nil) {
		(*callback)(int(posX), int(posY))
	}
}

