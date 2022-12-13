package web

import "testing"

func TestGetInstance(t *testing.T) {
	i1 := GetInstance()
	i2 := GetInstance()
	t.Logf("Test Web Single Instance:\ni1: %v, \ni2: %v", i1, i2)
	if i1 != i2 {
		t.Error("GetInstance() should return the same instance")
	}
}
