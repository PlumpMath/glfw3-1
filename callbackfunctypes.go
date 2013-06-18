package glfw3

type ErrorCallback func (int, string) //error, message

type IntPosCallback func (int, int) //x/width, y/height
type DoublePosCallback func (float64, float64) //x/width, y/height

type EventCallback func ()
type StatusCallback func (bool)

type KeyCallback func (int, int) //key, action
type CharCallback func (uint) //character
type MouseButtonCallback func (int, int) //button, action
