/*
Ogre3d bindings for Go. Uses llcoi to bind to ogre via c.
*/

package ogre

// #cgo LDFLAGS: -lllcoi
// #include "llcoi/ogre_interface.h"
// #include "llcoi/main.h"
import "C"

type Vector2 struct {
	cptr C.Vector2Handle
}

func CreateVector2() (Vector2) {
	var result Vector2
	result.cptr = C.vector2_create()	
	return result
}

func CreateVector2FromValues(x float32, y float32) (Vector2) {
	var result Vector2
	result.cptr = C.vector2_create_from_values(C.coiReal(x), C.coiReal(y))	
	return result
}
