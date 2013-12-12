package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type Entity struct {
	cptr C.EntityHandle
}
