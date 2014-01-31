package ogre

/*
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

func SetDefaultNumMipmaps(num int) {
	tm := C.texturemanager_singleton()
	C.texturemanager_set_default_num_mipmaps(tm, C.int(num))
}
