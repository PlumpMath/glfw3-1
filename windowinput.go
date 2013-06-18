package glfw3

//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
//#include "callback.h"
import "C"
import "unsafe"


func (window Window) GetInputMode(mode int) int {
	return (int)(C.glfwGetInputMode(window.toC(), C.int(mode)))
}

func (window Window) SetInputMode(mode, value int) {
	C.glfwSetInputMode(window.toC(), C.int(mode), C.int(value))
}

func (window Window) GetKey(key int) int {
	return (int)(C.glfwGetKey(window.toC(), C.int(key)))
}

func (window Window) GetMouseButton(button int) int {
	return (int)(C.glfwGetMouseButton(window.toC(), C.int(button)))
}

func (window Window) GetCursorPos() (float64, float64) {
	var xpos, ypos C.double
	C.glfwGetCursorPos(window.toC(), &xpos, &ypos)
	return float64(xpos), float64(ypos)
}

func (window Window) SetCursorPos(xpos, ypos float64) {
	C.glfwSetCursorPos(window.toC(), C.double(xpos), C.double(ypos))
}




func (w *Window) SetCursorEnterCallback(callback StatusCallback) {
	w.cursorEnterCallback = &callback
	C.cSetCursorEnterCallback(w.pointer)
}

func (w *Window) SetKeyCallback(callback KeyCallback) {
	w.keyCallback = &callback
	C.cSetKeyCallback(w.pointer)
}

func (w *Window) SetCharCallback(callback CharCallback) {
	w.charCallback = &callback
	C.cSetCharCallback(w.pointer)
}

func (w *Window) SetMouseButtonCallback(callback MouseButtonCallback) {
	w.mouseButtonCallback = &callback
	C.cSetMouseButtonCallback(w.pointer)
}

func (w *Window) SetCursorPosCallback(callback DoublePosCallback) {
	w.cursorPosCallback = &callback
	C.cSetCursorPosCallback(w.pointer)
}

func (w *Window) SetScrollCallback(callback DoublePosCallback) {
	w.scrollCallback = &callback
	C.cSetScrollCallback(w.pointer)
}



//export goCursorEnterCallback
func goCursorEnterCallback(window unsafe.Pointer, entered int) {
	callback := getGoWindow(unsafe.Pointer(window)).cursorEnterCallback
	if (callback != nil) {
		(*callback)(entered != 0)
	}
}

//export goKeyCallback
func goKeyCallback(window unsafe.Pointer, key, action int) {
	callback := getGoWindow(unsafe.Pointer(window)).keyCallback
	if (callback != nil) {
		(*callback)(key, action)
	}
}

//export goCharCallback
func goCharCallback(window unsafe.Pointer, character uint) {
	callback := getGoWindow(unsafe.Pointer(window)).charCallback
	if (callback != nil) {
		(*callback)(character)
	}
}

//export goMouseButtonCallback
func goMouseButtonCallback(window unsafe.Pointer, button, action int) {
	callback := getGoWindow(unsafe.Pointer(window)).mouseButtonCallback
	if (callback != nil) {
		(*callback)(button, action)
	}
}

//export goCursorPosCallback
func goCursorPosCallback(window unsafe.Pointer, posX, posY float64) {
	callback := getGoWindow(unsafe.Pointer(window)).cursorPosCallback
	if (callback != nil) {
		(*callback)(posX, posY)
	}
}

//export goScrollCallback
func goScrollCallback(window unsafe.Pointer, offsetX, offsetY float64) {
	callback := getGoWindow(unsafe.Pointer(window)).scrollCallback
	if (callback != nil) {
		(*callback)(offsetX, offsetY)
	}
}

