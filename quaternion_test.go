package ogre

import "fmt"
import "testing"

func TestQuaternionMultiplyQuaternion(t *testing.T) {
	quat := CreateQuaternionFromValues(1,2,3,4)
	fmt.Printf("Quaternion 1: %f %f %f %f\n", quat.W(), quat.X(), quat.Y(), quat.Z())
	quat2 := CreateQuaternionFromValues(1,2,3,4)
	fmt.Printf("Quaternion 2: %f %f %f %f\n", quat2.W(), quat2.X(), quat2.Y(), quat2.Z())
	quat.MultiplyQuaternion(quat2)
	fmt.Printf("Quaternion multiply result: %f %f %f %f\n", quat.W(), quat.X(), quat.Y(), quat.Z())	
	if quat.W() != -28 && quat.X() != 4 && quat.Y() != 6 && quat.Z() != 8 {
		t.Errorf("Error with multiplying quaternion")
	}
}

func TestQuaternionSubtractQuaternion(t *testing.T) {
	quat := CreateQuaternionFromValues(5,6,7,8)
	fmt.Printf("Quaternion 1: %f %f %f %f\n", quat.W(), quat.X(), quat.Y(), quat.Z())
	quat2 := CreateQuaternionFromValues(8,7,6,5)
	fmt.Printf("Quaternion 2: %f %f %f %f\n", quat2.W(), quat2.X(), quat2.Y(), quat2.Z())
	quat.SubtractQuaternion(quat2)
	fmt.Printf("Quaternion subtract result: %f %f %f %f\n", quat.W(), quat.X(), quat.Y(), quat.Z())	
	if quat.W() != -3 && quat.X() != -1 && quat.Y() != 1 && quat.Z() != 3 {
		t.Errorf("Error with subtracting quaternion")
	}
}
