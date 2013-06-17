package glfw3

//#cgo   linux  CFLAGS: -I/usr/local/include -pthread
//#cgo   linux LDFLAGS: -L/usr/local/lib -pthread -lX11 -lXrandr -lm -lGL -lrt -lglfw
//#cgo  darwin  CFLAGS: -I/usr/local/include
//#cgo  darwin LDFLAGS: -L/usr/local/lib -framework Cocoa -framework OpenGL -framework IOKit -lglfw3
//#cgo windows LDFLAGS: -lglu32 -lopengl32 -lglfwdll
//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GLFW/glfw3.h>
//#include "callback.h"
import "C"
import (
	"errors"
	"unsafe"
)

func Init() (err error) {
	if C.glfwInit() != 1 {
		err = errors.New("Failed to initialize GLFW")
	}
	return
}

func Terminate() {
	C.glfwTerminate()
}

func CreateWindow(width int, height int, title string, monitor *Monitor, share *Window) (window *Window, err error) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	windowPointer := (unsafe.Pointer)(C.glfwCreateWindow(
		C.int(width), C.int(height), ctitle,
		Cmonitor(monitor), Cwindow(share)))

	if (windowPointer != nil) {
		window = getGoWindow(windowPointer)
	} else {
		err = errors.New("Failed to open window")
	}

	return
}

func GetMonitors() ([]*Monitor, error) {
	var count C.int

	var cmonitor C.GLFWmonitor
	cmonitorSize := int(unsafe.Sizeof(cmonitor))

	cmonitorList := C.glfwGetMonitors(&count)

	if count == 0 {
		return nil, errors.New("An error occured")
	}

	list := make([]*Monitor, count)
	for i := range list {
		p := (*C.GLFWmonitor)(unsafe.Pointer(uintptr(unsafe.Pointer(cmonitorList)) + uintptr(i*cmonitorSize)))
		list[i] = &Monitor{(unsafe.Pointer)(p)}
	}

	return list, nil
}

func GetPrimaryMonitor() (monitor *Monitor, err error) {
	monitorPointer := (unsafe.Pointer)(C.glfwGetPrimaryMonitor())

	if (monitorPointer != nil) {
		monitor = &Monitor{monitorPointer}
	} else {
		err = errors.New("No window's context is current")
	}

	return
}

func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

func WindowHint(target, hint int) {
	C.glfwWindowHint(C.int(target), C.int(hint))
}

func GetVersion() (int, int, int) {
	var major, minor, rev C.int
	C.glfwGetVersion(&major, &minor, &rev)
	return int(major), int(minor), int(rev)
}

func GetVersionString() (string) {
	return C.GoString(C.glfwGetVersionString())
}

func GetTime() float64 {
	return float64(C.glfwGetTime())
}

func SetTime(time float64) {
	C.glfwSetTime(C.double(time))
}

func GetCurrentContext() (window *Window, err error) {
	windowPointer := (unsafe.Pointer)(C.glfwGetCurrentContext())

	if (windowPointer != nil) {
		window = getGoWindow(windowPointer)
	} else {
		err = errors.New("No window's context is current")
	}

	return
}

func SetSwapInterval(interval int) {
	C.glfwSwapInterval(C.int(interval))
}

func ExtensionSupported(name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.glfwExtensionSupported(cname) == 1
}

func ProcAddress(name string) unsafe.Pointer {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return (unsafe.Pointer)(C.glfwGetProcAddress(cname))
}

func PollEvents() {
	C.glfwPollEvents()
}

func WaitEvents() {
	C.glfwWaitEvents()
}


var errorCallback ErrorCallback
func SetErrorCallback(callback ErrorCallback) {
	errorCallback = callback
	C.cSetErrorCallback()
}
//export goErrorCallback
func goErrorCallback(err C.int, description *C.char) {
	if errorCallback != nil {
		errorCallback((int)(err), C.GoString(description))
	}
}

