package glfw3

//#cgo   linux  CFLAGS: -I/usr/local/include -pthread
//#cgo   linux LDFLAGS: -L/usr/local/lib -pthread -lX11 -lXrandr -lm -lGL -lrt -lglfw
//#cgo  darwin  CFLAGS: -I/usr/local/include
//#cgo  darwin LDFLAGS: -L/usr/local/lib -framework Cocoa -framework OpenGL -framework IOKit -lglfw
//#cgo windows LDFLAGS: -lglu32 -lopengl32 -lglfwdll
//#include <stdlib.h>
//#ifdef _WIN32
//  #define GLFW_DLL
//#endif
//#include <GL/glfw3.h>
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
		window = &Window{windowPointer}
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
		window = &Window{windowPointer}
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





/*

func Key(key int) int { return int(C.glfwGetKey(C.int(key))) }

func MouseButton(btn int) int { return int(C.glfwGetMouseButton(C.int(btn))) }

func MousePos() (int, int) {
	var cx, cy C.int
	C.glfwGetCursorPos(&cx, &cy)
	return int(cx), int(cy)
}

func SetMousePos(x, y int) { C.glfwSetCursorPos(C.int(x), C.int(y)) }

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
