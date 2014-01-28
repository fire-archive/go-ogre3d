/*
Ogre3d bindings for Go. Uses llcoi to bind to ogre via c.
*/

package ogre

// #cgo LDFLAGS: -lllcoi
// #include "llcoi/ogre_interface.h"
// #include "llcoi/main.h"
import "C"

type Vector3 struct {
	cptr C.Vector3Handle
}

func CreateVector3() (Vector3) {
	var result Vector3
	result.cptr = C.vector3_create()	
	return result
}

func (vec *Vector3) SetX(real float32) {
	C.vector3_set_x(vec.cptr, C.coiReal(real))
}

func (vec *Vector3) SetY(real float32) {
	C.vector3_set_y(vec.cptr, C.coiReal(real))
}

func (vec *Vector3) SetZ(real float32) {
	C.vector3_set_z(vec.cptr, C.coiReal(real))
}

func (vec *Vector3) Zero() {
	temp := C.vector3_ZERO()
	vec.SetX(float32(temp.x))
	vec.SetY(float32(temp.y))
	vec.SetZ(float32(temp.z))
}
