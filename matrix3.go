package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type Matrix3 struct {
	cmat C.coiMatrix3
}

func (mat *Matrix3) FromEulerAnglesYXZ(y, x, z float32) {
	mat.cmat = C.matrix_3_from_euler_angles_Y_X_Z(C.coiRadian(y), C.coiRadian(x), C.coiRadian(z))
}
