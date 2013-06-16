package glfw3

//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
import "C"


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


//SetKeyCallback
//SetCharCallback
//SetMouseButtonCallback
//SetCursorPosCallback
//SetCursorEnterCallback
//SetScrollCallback

