package glfw3

//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GL/glfw3.h>
import "C"
import (
	"unsafe"
)

type Monitor struct {
	pointer unsafe.Pointer
}

func (monitor Monitor) toC() (*C.GLFWmonitor) {
	return (*C.GLFWmonitor)(monitor.pointer)
}

func Cmonitor(monitor *Monitor) (cmonitor *C.GLFWmonitor) {
	if (monitor != nil) {
		cmonitor = monitor.toC()
	}
	return
}



//GetPos
//GetPhysicalSize
//GetName
//SetCallback
//GetVideoModes
//GetVideoMode

