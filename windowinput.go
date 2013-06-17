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

//export goCursorEnterCallback
func goCursorEnterCallback(window unsafe.Pointer, entered int) {
	callback := getGoWindow(unsafe.Pointer(window)).cursorEnterCallback
	if (callback != nil) {
		(*callback)(entered != 0)
	}
}

//TODO SetKeyCallback
//TODO SetCharCallback
//TODO SetMouseButtonCallback
//TODO SetCursorPosCallback
//TODO SetScrollCallback

