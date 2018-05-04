package main

import (
	"math"
	"testing"
)

var epsilon = 0.000001

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

// TestMatrix44fValues tests the 16 values of a Matrix44f.
func TestMatrix44fValues(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p float64, matrix Matrix44f, t *testing.T) {
	TestFloatValue("a", a, matrix.get(0, 0), t)
	TestFloatValue("b", b, matrix.get(0, 1), t)
	TestFloatValue("c", c, matrix.get(0, 2), t)
	TestFloatValue("d", d, matrix.get(0, 3), t)
	TestFloatValue("e", e, matrix.get(1, 0), t)
	TestFloatValue("f", f, matrix.get(1, 1), t)
	TestFloatValue("g", g, matrix.get(1, 2), t)
	TestFloatValue("h", h, matrix.get(1, 3), t)
	TestFloatValue("i", i, matrix.get(2, 0), t)
	TestFloatValue("j", j, matrix.get(2, 1), t)
	TestFloatValue("k", k, matrix.get(2, 2), t)
	TestFloatValue("l", l, matrix.get(2, 3), t)
	TestFloatValue("m", m, matrix.get(3, 0), t)
	TestFloatValue("n", n, matrix.get(3, 1), t)
	TestFloatValue("o", o, matrix.get(3, 2), t)
	TestFloatValue("p", p, matrix.get(3, 3), t)
}

// TestFloatValue compares a float value to an expected value.
func TestFloatValue(name string, expected, value float64, t *testing.T) {
	if math.Abs(expected-value) > epsilon {
		t.Errorf("%v should be %v, was %v, with epsilon %v", name, expected, value, epsilon)
	}
}

// TestBooleanValue compares a boolean value to an expected value.
func TestBooleanValue(name string, expected, value bool, t *testing.T) {
	if value != expected {
		t.Errorf("%v should be %v, was %v", name, expected, value)
	}
}
