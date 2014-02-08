/*
Ogre3d bindings for Go. Uses llcoi to bind to ogre via c.
*/

package ogre

// #cgo LDFLAGS: -lllcoi
// #include "llcoi/ogre_interface.h"
// #include "llcoi/main.h"
import "C"

type MovableObject struct {
	cptr C.MovableObjectHandle
}

// Get the base class of Entity.
func GetEntityBase(n Entity) MovableObject {
	var result MovableObject
	temp := n
	result.cptr = C.MovableObjectHandle(temp.cptr)
	return result
}

// Get the base class of Manual Object.
func GetManualObjectBase(o ManualObject) MovableObject {
	var result MovableObject
	temp := o
	result.cptr = C.MovableObjectHandle(temp.cptr)
	return result
}
