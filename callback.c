#include <GLFW/glfw3.h>
#include "callback.h"


//***********************************************
//GO Callback Functions
//***********************************************

extern void goErrorCallback(int, char*);
extern void goPosCallback(void* window, int posX, int posY);
extern void goSizeCallback(void* window, int width, int height);
extern void goCloseCallback(void* window);
extern void goRefreshCallback(void* window);
extern void goFocusCallback(void* window, int focused);
extern void goIconifyCallback(void* window, int iconified);
//extern void goFramebufferSizeCallback(void* window, int width, int height);
extern void goCursorEnterCallback(void* window, int entered);


//***********************************************
//C Callback Functions
//***********************************************

void cErrorCallback(int error, const char* description) {
	goErrorCallback(error, (char *)description);
}
void cPosCallback(GLFWwindow* window, int posX, int posY) {
	goPosCallback(window, posX, posY);
}
void cSizeCallback(GLFWwindow* window, int width, int height) {
	goSizeCallback(window, width, height);
}
void cCloseCallback(GLFWwindow* window) {
	goCloseCallback(window);
}
void cRefreshCallback(GLFWwindow* window) {
	goRefreshCallback(window);
}
void cFocusCallback(GLFWwindow* window, int focused) {
	goFocusCallback(window, focused);
}
void cIconifyCallback(GLFWwindow* window, int iconified) {
	goIconifyCallback(window, iconified);
}
/*
void cFramebufferSizeCallback(GLFWwindow* window, int width, int height) {
	goFramebufferSizeCallback(window, width, height);
}
*/
void cCursorEnterCallback(GLFWwindow* window, int entered) {
	goCursorEnterCallback(window, entered);
}


//***********************************************
//C SetCallback Functions
//***********************************************

void cSetErrorCallback(void) {
	glfwSetErrorCallback(cErrorCallback);
}
void cSetPosCallback(void* window) {
	glfwSetWindowPosCallback((GLFWwindow*)window, cPosCallback);
}
void cSetSizeCallback(void* window) {
	glfwSetWindowSizeCallback((GLFWwindow*)window, cSizeCallback);
}
void cSetCloseCallback(void* window) {
	glfwSetWindowCloseCallback((GLFWwindow*)window, cCloseCallback);
}
void cSetRefreshCallback(void* window) {
	glfwSetWindowRefreshCallback((GLFWwindow*)window, cRefreshCallback);
}
void cSetFocusCallback(void* window) {
	glfwSetWindowFocusCallback((GLFWwindow*)window, cFocusCallback);
}
void cSetIconifyCallback(void* window) {
	glfwSetWindowIconifyCallback((GLFWwindow*)window, cIconifyCallback);
}
/*
void cSetFramebufferSizeCallback(void* window) {
	glfwSetFramebufferSizeCallback((GLFWwindow*)window, cFramebufferSizeCallback);
}
*/
void cSetCursorEnterCallback(void* window) {
	glfwSetCursorEnterCallback((GLFWwindow*)window, cCursorEnterCallback);
}
