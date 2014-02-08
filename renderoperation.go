/*
Ogre3d bindings for Go. Uses llcoi to bind to ogre via c.
*/

package ogre

// #cgo LDFLAGS: -lllcoi
// #include "llcoi/ogre_interface.h"
// #include "llcoi/main.h"
import "C"

type OperationType int

const (
	OT_POINT_LIST             OperationType = 1
	OT_LINE_LIST              OperationType = 2
	OT_LINE_STRIP             OperationType = 3
	OT_TRIANGLE_LIST          OperationType = 4
	OT_TRIANGLE_STRIP         OperationType = 5
	OT_TRIANGLE_FAN           OperationType = 6
	OT_PATCH_1_CONTROL_POINT  OperationType = 7
	OT_PATCH_2_CONTROL_POINT  OperationType = 8
	OT_PATCH_3_CONTROL_POINT  OperationType = 9
	OT_PATCH_4_CONTROL_POINT  OperationType = 10
	OT_PATCH_5_CONTROL_POINT  OperationType = 11
	OT_PATCH_6_CONTROL_POINT  OperationType = 12
	OT_PATCH_7_CONTROL_POINT  OperationType = 13
	OT_PATCH_8_CONTROL_POINT  OperationType = 14
	OT_PATCH_9_CONTROL_POINT  OperationType = 15
	OT_PATCH_10_CONTROL_POINT OperationType = 16
	OT_PATCH_11_CONTROL_POINT OperationType = 17
	OT_PATCH_12_CONTROL_POINT OperationType = 18
	OT_PATCH_13_CONTROL_POINT OperationType = 19
	OT_PATCH_14_CONTROL_POINT OperationType = 20
	OT_PATCH_15_CONTROL_POINT OperationType = 21
	OT_PATCH_16_CONTROL_POINT OperationType = 22
	OT_PATCH_17_CONTROL_POINT OperationType = 23
	OT_PATCH_18_CONTROL_POINT OperationType = 24
	OT_PATCH_19_CONTROL_POINT OperationType = 25
	OT_PATCH_20_CONTROL_POINT OperationType = 26
	OT_PATCH_21_CONTROL_POINT OperationType = 27
	OT_PATCH_22_CONTROL_POINT OperationType = 28
	OT_PATCH_23_CONTROL_POINT OperationType = 29
	OT_PATCH_24_CONTROL_POINT OperationType = 30
	OT_PATCH_25_CONTROL_POINT OperationType = 31
	OT_PATCH_26_CONTROL_POINT OperationType = 32
	OT_PATCH_27_CONTROL_POINT OperationType = 33
	OT_PATCH_28_CONTROL_POINT OperationType = 34
	OT_PATCH_29_CONTROL_POINT OperationType = 35
	OT_PATCH_30_CONTROL_POINT OperationType = 36
	OT_PATCH_31_CONTROL_POINT OperationType = 37
	OT_PATCH_32_CONTROL_POINT OperationType = 38
)
