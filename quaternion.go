package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"
type Quaternion struct {
	cptr C.QuaternionHandle
}

func CreateQuaternion() Quaternion {
	var result Quaternion
	result.cptr = C.quaternion_create()
	return result
}

func (quat *Quaternion) FromRotationMatrix(mat Matrix3) {
	m := C.coiMatrix3(mat.cmat)
	C.quaternion_from_rotation_matrix(quat.cptr, &m)
}

func (quat *Quaternion) FromValues(w, x, y, z float32) Quaternion {
	var result Quaternion
	result.cptr = C.quaternion_from_values(C.coiReal(w), C.coiReal(x), C.coiReal(y), C.coiReal(z))
	return result	
}

func (quat *Quaternion) ToRotationMatrix(mat *Matrix3) {
	C.quaternion_to_rotation_matrix(quat.cptr, &mat.cmat)
}

// Get the w component of the quaternion.
func (quat *Quaternion) W() float32 {
	
	return float32(C.quaternion_get_w(quat.cptr))
}

// Get the x component of the quaternion.
func (quat *Quaternion) X() float32 {
	
	return float32(C.quaternion_get_x(quat.cptr))
}

// Get the y component of the quaternion.
func (quat *Quaternion) Y() float32 {
	
	return float32(C.quaternion_get_y(quat.cptr))
}

// Get the z component of the quaternion.
func (quat *Quaternion) Z() float32 {
	
	return float32(C.quaternion_get_z(quat.cptr))
}
