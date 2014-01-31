package ogre

import "fmt"
import "testing"

func TestFromEulerAnglesYXZ(t *testing.T) {
	var mat Matrix3
	mat.FromEulerAnglesYXZ(180, 180, 0)
	fmt.Printf("%d\n", mat.cmat)
}
