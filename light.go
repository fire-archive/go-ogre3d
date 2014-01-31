package ogre

/*
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type Light struct {
	cptr C.LightHandle
}

func (l *Light) SetPosition(x, y, z float32) {
	C.light_set_position(l.cptr, C.float(x), C.float(y), C.float(z))
}
