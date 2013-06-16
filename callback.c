#include <GLFW/glfw3.h>

extern void goErrorCallback(int, char*);
void cErrorCallback(int error, const char* description) {
	goErrorCallback(error, (char *)description);
}
void cSetErrorCallback (void) {
	glfwSetErrorCallback(cErrorCallback);
}

