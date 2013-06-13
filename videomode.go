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

type VideoMode struct {
	Width int
	Height int
	RedBits int
	GreenBits int
	BlueBits int
//	RefreshRate int
}

func videoModeFromC(ptr *C.GLFWvidmode) *VideoMode {
	return &VideoMode{
		int(ptr.width),
		int(ptr.height),
		int(ptr.redBits),
		int(ptr.greenBits),
		int(ptr.blueBits),
//		int(ptr.refreshRate),
	}
}

func (videoMode VideoMode) toC() (ptr *C.GLFWvidmode) {
	ptr = (*C.GLFWvidmode)(C.malloc(C.size_t(unsafe.Sizeof(ptr))))
	ptr.width = C.int(videoMode.Width)
	ptr.height = C.int(videoMode.Height)
	ptr.redBits = C.int(videoMode.RedBits)
	ptr.greenBits = C.int(videoMode.GreenBits)
	ptr.blueBits = C.int(videoMode.BlueBits)
//	ptr.refreshRate = C.int(videoMode.RefreshRate)
	return
}

