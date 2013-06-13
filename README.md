glfw3
=====

go bindings for glfw3

This library targets GLFW 3.0 and doesn't work with previous versions!!!
(http://www.glfw.org/download.html)

### Info

Missing functions:
* SetUserPointer and GetUserPointer
* SetGamma, GetGammaRamp and SetGammaRamp
* GetWindowAttrib
* GetFramebufferSize

Most of the callbacks are missing as well as most of the monitor and joystick stuff.

### Info

GLFW is licensed under the zlib/libpng license.
Its contents can be found in the LICENSE_GLFW file.

Some functions are ported from github.com/go-gl/glfw
Their license can be found in the LICENSE_GO-GL_GLFW file.


