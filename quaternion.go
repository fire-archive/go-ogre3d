package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type Quaternion struct {
	cquat C.coiQuaternion
	W, X, Y, Z float32
}

func (quat *Quaternion) FromRotationMatrix(mat Matrix3) {
	quat.cquat = C.quaternion_from_rotation_matrix(C.coiMatrix3(mat.cmat))
	quat.W = float32(quat.cquat.w)
	quat.X = float32(quat.cquat.x)
	quat.Y = float32(quat.cquat.y)
	quat.Z = float32(quat.cquat.z)
}

