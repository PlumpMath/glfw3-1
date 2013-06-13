package glfw3
/*
//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GL/glfw3.h>
import "C"
import (
	"fmt"
	"unsafe"
)
*/



//JoystickPresent
//GetJoystickAxes
//GetJoystickButtons
//GetJoystickName


/*

func JoystickParam(joy, param int) int { return int(C.glfwGetJoystickParam(C.int(joy), C.int(param))) }

func JoystickPos(joy int, axes []float32) int {
	if len(axes) == 0 {
		return 0
	}

	return int(C.glfwGetJoystickAxes(C.int(joy), (*C.float)(unsafe.Pointer(&axes[0])), C.int(len(axes))))
}

func JoystickButtons(joy int, buttons []byte) int {
	if len(buttons) == 0 {
		return 0
	}

	return int(C.glfwGetJoystickButtons(C.int(joy),
		(*C.uchar)(unsafe.Pointer(&buttons[0])), C.int(len(buttons))))
}

*/
