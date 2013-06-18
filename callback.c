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
extern void goKeyCallback(void* window, int key, int action);
extern void goCharCallback(void* window, int character);
extern void goMouseButtonCallback(void* window, int button, int action);
extern void goCursorPosCallback(void* window, int posX, int posY);
extern void goScrollCallback(void* window, int offsetX, int offsetY);


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
void cKeyCallback(GLFWwindow* window, int key, int action) {
	goKeyCallback(window, key, action);
}
void cCharCallback(GLFWwindow* window, unsigned int character) {
	goCharCallback(window, character);
}
void cMouseButtonCallback(GLFWwindow* window, int button, int action) {
	goMouseButtonCallback(window, button, action);
}
void cCursorPosCallback(GLFWwindow* window, double posX, double posY) {
	goCursorPosCallback(window, posX, posY);
}
void cScrollCallback(GLFWwindow* window, double offsetX, double offsetY) {
	goScrollCallback(window, offsetX, offsetY);
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
void cSetKeyCallback(void* window) {
	glfwSetKeyCallback((GLFWwindow*)window, cKeyCallback);
}
void cSetCharCallback(void* window) {
	glfwSetCharCallback((GLFWwindow*)window, cCharCallback);
}
void cSetMouseButtonCallback(void* window) {
	glfwSetMouseButtonCallback((GLFWwindow*)window, cMouseButtonCallback);
}
void cSetCursorPosCallback(void* window) {
	glfwSetCursorPosCallback((GLFWwindow*)window, cCursorPosCallback);
}
void cSetScrollCallback(void* window) {
	glfwSetScrollCallback((GLFWwindow*)window, cScrollCallback);
}

