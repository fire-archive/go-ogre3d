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
	mat.cmat = C.matrix_3_from_euler_angles_y_x_z(C.coiRadian(y), C.coiRadian(x), C.coiRadian(z))
}

func (mat *Matrix3) ToEulerAnglesYXZ(y, x, z *float32) {
	var yangle, xangle, zangle C.coiRadian
	C.matrix_3_to_euler_angles_y_x_z(mat.cmat, &yangle, &xangle, &zangle)
	*y = float32(yangle)
	*x = float32(xangle)
	*z = float32(zangle)
}
