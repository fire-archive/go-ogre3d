package ogre

/*
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ois_interface.h"
*/
import "C"

type InputManager struct {
	cptr C.InputManagerHandle
}

func NewInput(w RenderWindow) InputManager {
	var result InputManager
	result.cptr = C.create_input_system(C.uint(w.Handle()))
	return result
}

func (i *InputManager) Destroy() {
	C.destroy_input_system(i.cptr)
	i.cptr = nil
}

func (i *InputManager) NewKeyboard(buffered bool) Keyboard {
	var result Keyboard
	result.cptr = C.create_keyboard_object(i.cptr, cbool(buffered))
	return result
}

func (i *InputManager) DestroyKeyboard(k Keyboard) {
	C.destroy_keyboard_object(i.cptr, k.cptr)
	k.cptr = nil
}

func (i *InputManager) NewMouse(buffered bool) Mouse {
	var result Mouse
	result.cptr = C.create_mouse_object(i.cptr, cbool(buffered))
	return result
}

func (i *InputManager) DestroyMouse(m Mouse) {
	C.destroy_mouse_object(i.cptr, m.cptr)
	m.cptr = nil
}
