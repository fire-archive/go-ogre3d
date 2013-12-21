package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type RenderSystem struct {
	cptr C.RenderSystemHandle
}

type RenderSystemList struct {
	cptr C.RenderSystemListHandle
}

func (r *RenderSystem) SetConfigOption(key, value string) {
	C.render_system_set_config_option(r.cptr, C.CString(key), C.CString(value))
}

func (r *RenderSystemList) RenderSystemListSize() uint {
	return uint(C.render_system_list_size(r.cptr))
}
