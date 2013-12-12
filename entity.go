package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #cgo CFLAGS: -I./llcoihdr
 #include "llcoi/ogre_interface.h"
*/
import "C"

type Entity struct {
	cptr C.EntityHandle
}
