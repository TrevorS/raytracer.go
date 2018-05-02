package main

import "testing"

// TestVec2fValues tests the 2 members of a Vec2f.
func TestVec2fValues(x, y float64, v Vec2f, t *testing.T) {
	TestFloatValue("x", x, v.x, t)
	TestFloatValue("y", y, v.y, t)
}

// TestVec3fValues tests the 3 members of a Vec3f.
func TestVec3fValues(x, y, z float64, v Vec3f, t *testing.T) {
	TestFloatValue("x", x, v.x, t)
	TestFloatValue("y", y, v.y, t)
	TestFloatValue("z", z, v.z, t)
}

// TestFloatValue compares a float value to an expected value.
func TestFloatValue(name string, expected, value float64, t *testing.T) {
	if value != expected {
		t.Errorf("%v should be %v, was %v", name, expected, value)
	}
}
