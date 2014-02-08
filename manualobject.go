/*
Ogre3d bindings for Go. Uses llcoi to bind to ogre via c.
*/

package ogre

// #cgo LDFLAGS: -lllcoi
// #include "llcoi/ogre_interface.h"
// #include "llcoi/main.h"
import "C"

type ManualObject struct {
	cptr C.ManualObjectHandle
}

type ManualObjectSection struct {
	cptr C.ManualObjectSectionHandle
}

func (o *ManualObject) SetDynamic(dyn bool) {
	C.manualobject_set_dynamic(o.cptr, cbool(dyn))
}

func (o *ManualObject) Begin(materialName string, opType OperationType, groupName string) {
	C.manualobject_begin(o.cptr, C.CString(materialName), C.operation_type(opType), C.CString(groupName))
}

func (o *ManualObject) BeginUpdate(sectionIndex uint) {
	C.manualobject_begin_update(o.cptr, C.size_t(sectionIndex))
}

func (o *ManualObject) Position(x, y, z float32) {
	var v C.coiVector3
	v.x = C.coiReal(x)
	v.y = C.coiReal(y)
	v.z = C.coiReal(z)
	C.manualobject_position(o.cptr, &v)
}

func (o *ManualObject) End() ManualObjectSection {
	var sectionObj ManualObjectSection
	sectionObj.cptr = C.manualobject_end(o.cptr)
	return sectionObj
}
