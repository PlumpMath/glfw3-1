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

extern void goSizeCallback(void* window, int width, int height);
void cSizeCallback(GLFWwindow* window, int width, int height) {
	goSizeCallback(window, width, height);
}
void cSetSizeCallback(void* window) {
	glfwSetWindowSizeCallback((GLFWwindow*)window, cSizeCallback);
}

extern void goCloseCallback(void* window);
void cCloseCallback(GLFWwindow* window) {
	goCloseCallback(window);
}
void cSetCloseCallback(void* window) {
	glfwSetWindowCloseCallback((GLFWwindow*)window, cCloseCallback);
}

extern void goRefreshCallback(void* window);
void cRefreshCallback(GLFWwindow* window) {
	goRefreshCallback(window);
}
void cSetRefreshCallback(void* window) {
	glfwSetWindowRefreshCallback((GLFWwindow*)window, cRefreshCallback);
}

extern void goFocusCallback(void* window, int focused);
void cFocusCallback(GLFWwindow* window, int focused) {
	goFocusCallback(window, focused);
}
void cSetFocusCallback(void* window) {
	glfwSetWindowFocusCallback((GLFWwindow*)window, cFocusCallback);
}

extern void goIconifyCallback(void* window, int iconified);
void cIconifyCallback(GLFWwindow* window, int iconified) {
	goIconifyCallback(window, iconified);
}
void cSetIconifyCallback(void* window) {
	glfwSetWindowIconifyCallback((GLFWwindow*)window, cIconifyCallback);
}

/*
extern void goFramebufferSizeCallback(void* window, int width, int height);
void cFramebufferSizeCallback(GLFWwindow* window, int width, int height) {
	goFramebufferSizeCallback(window, width, height);
}
void cSetFramebufferSizeCallback(void* window) {
	glfwSetFramebufferSizeCallback((GLFWwindow*)window, cFramebufferSizeCallback);
}
*/

