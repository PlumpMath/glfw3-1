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


/*

/*
func VideoModes(max int) []*VidMode {
	var vm C.GLFWvidmode

	size := int(unsafe.Sizeof(vm))
	ptr := (*C.GLFWvidmode)(C.malloc(C.size_t(size * max)))
	defer C.free(unsafe.Pointer(ptr))
	count := C.glfwGetVideoModes(ptr, C.int(max))

	if count == 0 {
		return nil
	}

	list := make([]*VidMode, count)
	for i := range list {
		p := (*C.GLFWvidmode)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i*size)))
		list[i] = vidModeFromPtr(p)
	}

	return list
}

func DesktopMode() *VidMode {
	var vm C.GLFWvidmode
	C.glfwGetVideoMode(&vm)
	return vidModeFromPtr(&vm)
}
*/

