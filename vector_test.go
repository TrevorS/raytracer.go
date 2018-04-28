package main

import (
	"math"
	"testing"
)

func testVectorValues(x, y, z float64, v Vector, t *testing.T) {
	if v.x != x {
		t.Errorf("x should be %v, was %v", x, v.x)
	}

	if v.y != y {
		t.Errorf("y should be %v, was %v", y, v.y)
	}

	if v.z != z {
		t.Errorf("z should be %v, was %v", z, v.z)
	}
}

func testValue(expected, value float64, t *testing.T) {
	if value != expected {
		t.Errorf("value should be %v, was %v", expected, value)
	}
}

func TestVector(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}

	testVectorValues(x, y, z, v, t)
}

func TestVectorZero(t *testing.T) {
	expected := 0.0

	v := VectorZero()

	testVectorValues(expected, expected, expected, v, t)
}

func TestVectorAt(t *testing.T) {
	expected := 3.0

	v := VectorAt(expected)

	testVectorValues(expected, expected, expected, v, t)
}

func TestVectorNegate(t *testing.T) {
	value := 1.0
	expected := -1 * value

	v := VectorAt(value).negate()

	testVectorValues(expected, expected, expected, v, t)
}

func TestVectorAdd(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	v2 := Vector{x, y, z}

	testVectorValues(
		x+x,
		y+y,
		z+z,
		v.add(v2),
		t,
	)
}

func TestVectorSubtract(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	v2 := Vector{x, y, z}

	testVectorValues(
		x-x,
		y-y,
		z-z,
		v.subtract(v2),
		t,
	)
}

func TestVectorMultiply(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	v2 := Vector{x, y, z}

	testVectorValues(
		x*x,
		y*y,
		z*z,
		v.multiply(v2),
		t,
	)
}

func TestVectorMultiplyScalar(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0
	value := 2.0

	v := Vector{x, y, z}

	testVectorValues(
		x*value,
		y*value,
		z*value,
		v.multiplyScalar(value),
		t,
	)
}

func TestVectorDotProduct(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	v2 := Vector{x, y, z}

	dotProduct := v.dotProduct(v2)
	expected := v.x*v2.x + v.y*v2.y + v.z*v2.z

	testValue(expected, dotProduct, t)
}

func TestVectorDivideScalar(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	value := 2.0

	testVectorValues(
		x/value,
		y/value,
		z/value,
		v.divideScalar(value),
		t,
	)
}

func TestVectorCrossProduct(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	v2 := Vector{x, y, z}

	testVectorValues(
		v.y*v2.z-v.z*v2.y,
		v.z*v2.x-v.x*v2.z,
		v.x*v2.y-v.y*v2.x,
		v.crossProduct(v2),
		t,
	)
}

func TestVectorNorm(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	norm := v.norm()

	expected := v.x*v.x + v.y*v.y + v.z*v.z

	testValue(expected, norm, t)
}

func TestVectorLength(t *testing.T) {
	x := 1.0
	y := 2.0
	z := 3.0

	v := Vector{x, y, z}
	expected := math.Sqrt(v.norm())

	testValue(expected, v.length(), t)
}
