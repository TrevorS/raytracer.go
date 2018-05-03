package main

import (
	"testing"
)

func TestVec2fZero(t *testing.T) {
	expected := 0.0

	v := Vec2fZero()

	TestVec2fValues(expected, expected, v, t)
}

func TestVec2fAt(t *testing.T) {
	expected := 3.0

	v := Vec2fAt(expected)

	TestVec2fValues(expected, expected, v, t)
}

func TestVec2fAdd(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	v2 := Vec2f{x, y}

	TestVec2fValues(
		x+x,
		y+y,
		v.add(v2),
		t,
	)
}

func TestVec2fSubtract(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	v2 := Vec2f{x, y}

	TestVec2fValues(
		x-x,
		y-y,
		v.subtract(v2),
		t,
	)
}

func TestVec2fMultiply(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	v2 := Vec2f{x, y}

	TestVec2fValues(
		x*x,
		y*y,
		v.multiply(v2),
		t,
	)
}

func TestVec2fDivide(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	v2 := Vec2f{x, y}

	TestVec2fValues(
		x/x,
		y/y,
		v.divide(v2),
		t,
	)
}

func TestVec2fMultiplyScalar(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	value := 2.0

	TestVec2fValues(
		x*value,
		y*value,
		v.multiplyScalar(value),
		t,
	)
}

func TestVec2fDivideScalar(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	value := 2.0

	TestVec2fValues(
		x/value,
		y/value,
		v.divideScalar(value),
		t,
	)
}

func TestVec2fMultiplyInPlace(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	v2 := Vec2f{x, y}

	v.multiplyInPlace(v2)

	TestVec2fValues(
		x*x,
		y*y,
		v,
		t,
	)
}

func TestVec2fDivideInPlace(t *testing.T) {
	x := 1.0
	y := 2.0

	v := Vec2f{x, y}
	v2 := Vec2f{x, y}

	v.divideInPlace(v2)

	TestVec2fValues(
		x/x,
		y/y,
		v,
		t,
	)
}
