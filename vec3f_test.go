package main

import (
	"math"
	"testing"
)

func TestVec3fZero(t *testing.T) {
	expected := 0.0

	v := Vec3fZero()

	TestVec3fValues(expected, expected, expected, v, t)
}

func TestVec3fAt(t *testing.T) {
	expected := 3.0

	v := Vec3fAt(expected)

	TestVec3fValues(expected, expected, expected, v, t)
}

func TestVec3fNegate(t *testing.T) {
	value := 1.0
	expected := -1 * value

	v := Vec3fAt(value).negate()

	TestVec3fValues(expected, expected, expected, v, t)
}

func TestVec3fAdd(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	v2 := Vec3f{x, y, z}

	TestVec3fValues(
		x+x,
		y+y,
		z+z,
		v.add(v2),
		t,
	)
}

func TestVec3fSubtract(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	v2 := Vec3f{x, y, z}

	TestVec3fValues(
		x-x,
		y-y,
		z-z,
		v.subtract(v2),
		t,
	)
}

func TestVec3fMultiply(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	v2 := Vec3f{x, y, z}

	TestVec3fValues(
		x*x,
		y*y,
		z*z,
		v.multiply(v2),
		t,
	)
}

func TestVec3fMultiplyScalar(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0
	value := 2.0

	v := Vec3f{x, y, z}

	TestVec3fValues(
		x*value,
		y*value,
		z*value,
		v.multiplyScalar(value),
		t,
	)
}

func TestVec3fDotProduct(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	v2 := Vec3f{x, y, z}

	dotProduct := v.dotProduct(v2)
	expected := v.x*v2.x + v.y*v2.y + v.z*v2.z

	TestFloatValue("dotProduct", expected, dotProduct, t)
}

func TestVec3fDivideScalar(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	value := 2.0

	TestVec3fValues(
		x/value,
		y/value,
		z/value,
		v.divideScalar(value),
		t,
	)
}

func TestVec3fCrossProduct(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	v2 := Vec3f{x, y, z}

	TestVec3fValues(
		v.y*v2.z-v.z*v2.y,
		v.z*v2.x-v.x*v2.z,
		v.x*v2.y-v.y*v2.x,
		v.crossProduct(v2),
		t,
	)
}

func TestVec3fNorm(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	norm := v.norm()

	expected := v.x*v.x + v.y*v.y + v.z*v.z

	TestFloatValue("norm", expected, norm, t)
}

func TestVec3fLength(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vec3f{x, y, z}
	expected := math.Sqrt(v.norm())

	TestFloatValue("length", expected, v.length(), t)
}
