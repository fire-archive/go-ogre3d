package ogre

/*
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type RenderWindow struct {
	cptr C.RenderWindowHandle
}

func (rw *RenderWindow) AddViewport(c Camera) Viewport {
	var result Viewport
	result.cptr = C.render_window_add_viewport(rw.cptr, c.cptr, 0, 0, 0, 1, 1)

	return result
}

func (rw *RenderWindow) IsClosed() bool {
	result := C.render_window_closed(rw.cptr)
	return gobool(result)
}

func (rw *RenderWindow) Handle() uint {
	return uint(C.render_window_get_hwnd(rw.cptr))
}

func (rw *RenderWindow) Reposition(left, top int) {
	C.render_window_reposition(rw.cptr, C.int(left), C.int(top))
}

func (rw *RenderWindow) SetVisible(visible bool) {
	C.render_window_set_visible(rw.cptr, cbool(visible))
}

func MessagePump() {
	C.pump_messages()
}
