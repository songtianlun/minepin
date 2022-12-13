package utils

import "testing"

func TestGetTypeString(t *testing.T) {
	if GetTypeString(1) != "int" {
		t.Errorf("GetTypeString(1) = %s; want int", GetTypeString(1))
	}
	if GetTypeString("1") != "string" {
		t.Errorf("GetTypeString(\"1\") = %s; want string", GetTypeString("1"))
	}
	if GetTypeString(1.1) != "float64" {
		t.Errorf("GetTypeString(1.1) = %s; want float64", GetTypeString(1.1))
	}
	if GetTypeString(true) != "bool" {
		t.Errorf("GetTypeString(true) = %s; want bool", GetTypeString(true))
	}
}
