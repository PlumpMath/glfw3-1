package glfw3

//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
import "C"
import (
	"unsafe"
	"errors"
)

type Window struct {
	pointer unsafe.Pointer
	posCallback *PosCallback
}

func (window Window) toC() (*C.GLFWwindow) {
	return (*C.GLFWwindow)(window.pointer)
}

func Cwindow(window *Window) (cwindow *C.GLFWwindow) {
	if (window != nil) {
		cwindow = window.toC()
	}
	return
}

var cWindowMap map[unsafe.Pointer]*Window = make(map[unsafe.Pointer]*Window)

func getGoWindow(cwindow unsafe.Pointer) (window *Window) {
	window, ok := cWindowMap[cwindow]
	if !ok {
		window = new(Window)
		window.pointer = cwindow
		cWindowMap[cwindow] = window
	}
	return
}



func (window Window) Destroy() {
	C.glfwDestroyWindow(window.toC())
}

func (window Window) ShouldClose() bool {
	return (int)(C.glfwWindowShouldClose(window.toC())) != 0
}

func (window Window) SetShouldClose(value bool) {
	if value {
		C.glfwSetWindowShouldClose(window.toC(), C.int(1))
	} else {
		C.glfwSetWindowShouldClose(window.toC(), C.int(0))
	}
}

func (window Window) SetTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	C.glfwSetWindowTitle(window.toC(), ctitle)
}

func (window Window) GetPos() (int, int) {
	var x, y C.int
	C.glfwGetWindowSize(window.toC(), &x, &y)
	return int(x), int(y)
}

func (window Window) SetPos(x, y int) {
	C.glfwSetWindowPos(window.toC(), C.int(x), C.int(y))
}

func (window Window) GetSize() (int, int) {
	var width, height C.int
	C.glfwGetWindowSize(window.toC(), &width, &height)
	return int(width), int(height)
}

func (window Window) SetSize(width int, height int) {
	C.glfwSetWindowSize(window.toC(), C.int(width), C.int(height))
}
/*
func (window Window) GetFramebufferSize() (int, int) {
	var width, height C.int
	C.glfwGetFramebufferSize(window.toC(), &width, &height)
	return int(width), int(height)
}
*/
func (window Window) Iconify() {
	C.glfwIconifyWindow(window.toC())
}

func (window Window) Restore() {
	C.glfwRestoreWindow(window.toC())
}

func (window Window) Show() {
	C.glfwShowWindow(window.toC())
}

func (window Window) Hide() {
	C.glfwHideWindow(window.toC())
}

func (window Window) GetMonitor() (monitor *Monitor, err error) {
	monitorPointer := (unsafe.Pointer)(C.glfwGetWindowMonitor(window.toC()))
	if (monitorPointer != nil) {
		monitor = &Monitor{monitorPointer}
	} else {
		err = errors.New("Window is in windowed mode")
	}
	return
}

/*
func (window Window) GetAttrib(attrib int) {
	return (int)(C.glfwGetWindowAttrib(window.toC(), C.int(attrib)))
}
*/

func (window Window) SetClipboardString(clipString string) {
	cclipString := C.CString(clipString)
	defer C.free(unsafe.Pointer(cclipString))
	C.glfwSetClipboardString(window.toC(), cclipString)
}

func (window Window) GetClipboardString() (string) {
	return C.GoString(C.glfwGetClipboardString(window.toC()))
}



func (window Window) MakeContextCurrent() {
	C.glfwMakeContextCurrent(window.toC())
}

func (window Window) SwapBuffers() {
	C.glfwSwapBuffers(window.toC())
}

