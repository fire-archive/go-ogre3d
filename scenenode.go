package ogre

/*
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type SceneNode struct {
	cptr C.SceneNodeHandle
}

func (n *SceneNode) Attach(e Entity) {
	C.scenenode_attach_entity(n.cptr, e.cptr)
}

func (n *SceneNode) AttachObject(o MovableObject) {
	C.scenenode_attach_object(n.cptr, o.cptr)
}

func (n *SceneNode) Detach(e Entity) {
	C.scenenode_detach_entity(n.cptr, e.cptr)
}

func (n *SceneNode) SetPosition(x, y, z float32) {
	C.scenenode_set_position(n.cptr, C.float(x), C.float(y), C.float(z))
}

func (n *SceneNode) SetVisible(visible bool) {
	C.scenenode_set_visible(n.cptr, cbool(visible))
}

type TransformSpace int

const (
	TS_LOCAL TransformSpace = iota
	TS_PARENT
	TS_WORLD
)

func (n *SceneNode) YawDeg(angle float32, t TransformSpace) {
	C.scenenode_yaw_degree(n.cptr, C.coiReal(angle), C.transform_space(t))
}

func (n *SceneNode) Yaw(radians float32, t TransformSpace) {
	C.scenenode_yaw(n.cptr, C.coiReal(radians), C.transform_space(t))
}

func (n *SceneNode) Pitch(radians float32, t TransformSpace) {
	C.scenenode_pitch(n.cptr, C.coiReal(radians), C.transform_space(t))
}

func (n *SceneNode) Roll(radians float32, t TransformSpace) {
	C.scenenode_roll(n.cptr, C.coiReal(radians), C.transform_space(t))
}

func (n *SceneNode) ResetOrientation() {
	C.scenenode_reset_orientation(n.cptr)
}

func (n *SceneNode) CreateChildSceneNode(name string, translate Vector3, rotate Quaternion) SceneNode {
	var result SceneNode
	result.cptr = C.scenenode_create_child_scenenode(n.cptr, C.CString(name), translate.cptr, rotate.cptr)
	return result
}

func (n *SceneNode) SetOrientation(rot Quaternion) {
	C.scenenode_set_orientation(n.cptr, C.float(rot.W()), C.float(rot.X()), C.float(rot.Y()), C.float(rot.Z()))
} 
