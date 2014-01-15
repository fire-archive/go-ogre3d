package ogre

/* 
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type Quaternion struct {
	cquat C.coiQuaternion
	w, x, y, z float32
}

func FromRotationMatrix(mat Matrix3) (Quaternion) {
	var quat Quaternion
	quat.cquat = C.quaternion_from_rotation_matrix(C.coiMatrix3(mat.cmat))
	quat.w = float32(quat.cquat.w)
	quat.x = float32(quat.cquat.x)
	quat.y = float32(quat.cquat.y)
	quat.z = float32(quat.cquat.z)
	return quat
}

