package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #cgo CFLAGS: -I./llcoihdr
 #include "llcoi/ogre_interface.h"
*/
import "C"

type RenderSystem struct {
	cptr C.RenderSystemHandle
}

func (r *RenderSystem) SetConfigOption(key, value string) {
	C.render_system_set_config_option(r.cptr, C.CString(key), C.CString(value))
}
