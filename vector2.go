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

func (vec *Vector2) MultiplyVector2(rhs Vector2) (Vector2) {
	var result Vector2
	result.cptr = C.vector2_multiply_vector2(vec.cptr, rhs.cptr)
	return result
}

func (vec *Vector2) MultiplyScalar(real float32) (Vector2) {
	var result Vector2
	result.cptr = C.vector2_multiply_scalar(vec.cptr, C.coiReal(real))
	return result
}

func (vec *Vector2) X() (float32) {
	return float32(C.vector2_x(vec.cptr))
}

func (vec *Vector2) Y() (float32) {
	return float32(C.vector2_y(vec.cptr))
}

func (vec *Vector2) SetX(real float32) {
	C.vector2_set_x(vec.cptr, C.coiReal(real))
}

func (vec *Vector2) SetY(real float32) {
	C.vector2_set_y(vec.cptr, C.coiReal(real))
}
