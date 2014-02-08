package ogre

/*
 #cgo LDFLAGS: -lllcoi
 #include "llcoi/ogre_interface.h"
*/
import "C"

type ResourceGroupManager struct {
	cptr C.ResourceGroupManagerHandle
}

func GetResourceGroupManager() ResourceGroupManager {
	var result ResourceGroupManager
	result.cptr = C.resourcegroupmanager_singleton()
	return result
}

func (rgm *ResourceGroupManager) AddResourceLocation(location, locationType, group string) {
	C.resourcegroupmanager_add_resource_location(rgm.cptr, C.CString(location), C.CString(locationType), C.CString(group))
}

func (rgm *ResourceGroupManager) InitialiseAllResourceGroups() {
	C.resourcegroupmanager_initialise_all_resourcegroups(rgm.cptr)
}

func DEFAULT_RESOURCE_GROUP_NAME() string {
	return C.GoString(C.resourcegroupmanager_DEFAULT_RESOURCE_GROUP_NAME())
}

func INTERNAL_RESOURCE_GROUP_NAME() string {
	return C.GoString(C.resourcegroupmanager_INTERNAL_RESOURCE_GROUP_NAME())
}

func AUTODETECT_RESOURCE_GROUP_NAME() string {
	return C.GoString(C.resourcegroupmanager_AUTODETECT_RESOURCE_GROUP_NAME())
}
