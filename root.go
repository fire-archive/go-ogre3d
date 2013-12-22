/*
Ogre3d bindings for Go. Uses llcoi to bind to ogre via c.
*/

package ogre

// #cgo LDFLAGS: -lllcoi
// #include "llcoi/ogre_interface.h"
// #include "llcoi/main.h"
import "C"

type Root struct {
	cptr C.RootHandle
}

type NameValuePairList struct {
	cptr C.NameValuePairListHandle
}

func createNameValuePairList() NameValuePairList {
	var name NameValuePairList
	name.cptr = C.create_name_value_pair_list()
	return name
}

func (n *NameValuePairList) addPair(name string, value string) {
	C.add_pair(n.cptr, C.CString(name), C.CString(value))
}

func (n *NameValuePairList) destroyNameValuePairList() {
	C.destroy_name_value_pair_list(n.cptr)
	n.cptr = nil
}

func NewRoot(pluginsCfg, ogreCfg, logfile string) Root {
	var result Root
	result.cptr = C.create_root(C.CString(pluginsCfg), C.CString(ogreCfg), C.CString(logfile))

	return result
}

func (r *Root)Destroy() {
	C.delete_root(r.cptr)
	r.cptr = nil
}

func RootInstance() Root {
	var result Root
	result.cptr = C.root_singleton()
	
	return result
}

func (r *Root) Initialised() bool {
	return gobool(C.root_is_initialised(r.cptr))
}

func (r *Root) SaveConfig() {
	C.save_config(r.cptr)
}

func (r *Root) RestoreConfig() bool {
	return gobool(C.restore_config(r.cptr))
}

func (r *Root) GetRenderSystemByName(name string) RenderSystem {
	var result RenderSystem
	result.cptr = C.get_render_system_by_name(r.cptr, C.CString(name))

	return result
}

func (r *Root) GetAvailableRenderers() RenderSystemList {
	var result RenderSystemList
	result.cptr = C.root_get_available_renderers(r.cptr)
	
	return result
}

func (r *Root) SetRenderSystem(s RenderSystem) {
	C.set_render_system(r.cptr, s.cptr)
}

func (r *Root) ShowConfigDialog() bool {
	result := C.show_config_dialog(r.cptr)
	return gobool(result)
}

func (r *Root) Initialise(createWindow bool, windowTitle string) RenderWindow {
	var result RenderWindow
	result.cptr = C.root_initialise(r.cptr, cbool(createWindow), C.CString(windowTitle))

	return result
}

func (r *Root) CreateRenderWindow(name string, width int, height int, fullScreen bool, params NameValuePairList) RenderWindow{
	var result RenderWindow
	result.cptr = C.create_render_window(r.cptr, C.CString(name), C.int(width), C.int(height), cbool(fullScreen), params.cptr)
	return result
}

func (r *Root) CreateSceneManager(typename, instancename string) SceneManager {
	var result SceneManager
	result.cptr = C.create_scene_manager(r.cptr, C.CString(typename), C.CString(instancename))

	return result
}

func (r *Root) RenderOneFrameEx(timestep float32) bool {
	return gobool(C.render_one_frame_ex(r.cptr, C.float(timestep)))
}

func (r *Root) RenderOneFrame() bool {
	return gobool(C.render_one_frame(r.cptr))
}

func (r *Root) LoadPlugin(name string) {
	C.load_ogre_plugin(r.cptr, C.CString(name))
}


