package glfw3

type ErrorCallback func (int, string) //error, message

type PosCallback func (int, int) //x, y
type SizeCallback func (int, int) //width, height
type EventCallback func ()
type StatusCallback func (bool)

type KeyCallback func (int, int, int, int) //key, scancode, action, mods
type CharCallback func (int) //character
type MouseButtonCallback func (int, int, int) //button, action, mods
