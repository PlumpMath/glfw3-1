#include <GLFW/glfw3.h>
#include "callback.h"
#include <stdio.h>

extern void goErrorCallback(int, char*);
void cErrorCallback(int error, const char* description) {
	goErrorCallback(error, (char *)description);

}
void cSetErrorCallback(void) {
	glfwSetErrorCallback(cErrorCallback);
}


extern void goPosCallback(void* window, int posX, int posY);
void cPosCallback(GLFWwindow* window, int posX, int posY) {
	goPosCallback(window, posX, posY);
}
void cSetPosCallback(void* window) {
	glfwSetWindowPosCallback((GLFWwindow*)window, cPosCallback);
}
