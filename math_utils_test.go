package main

import (
	"testing"
)

func TestClampHigh(t *testing.T) {
	low := 2.0
	high := 4.0
	value := 6.0

	result := Clamp(low, high, value)

	TestFloatValue("clamp high", high, result, t)
}

func TestClampLow(t *testing.T) {
	low := 2.0
	high := 4.0
	value := 1.0

	result := Clamp(low, high, value)

	TestFloatValue("clamp low", low, result, t)
}

func TestDegreesToRadians(t *testing.T) {
	degrees := 45.0
	radians := 0.7853981633974483

	result := DegreesToRadians(degrees)

	TestFloatValue("degrees to radians", radians, result, t)
}

func TestMix(t *testing.T) {
	v1 := Vec3f{1, 2, 3}
	v2 := Vec3f{2, 3, 4}
	mixValue := 2.0

	result := Mix(v1, v2, mixValue)

	TestVec3fValues(
		3,
		4,
		5,
		result,
		t,
	)
}

func TestSolveQuadraticWithTwoRoots(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(5, 6, 1)

	x0Expected := -1.0
	x1Expected := -0.2

	TestBooleanValue("hasRoots", true, hasRoots, t)

	TestFloatValue("x0", x0Expected, x0, t)
	TestFloatValue("x1", x1Expected, x1, t)
}

func TestSolveQuadraticWithTwoRootsNegativeB(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(5, -6, 1)

	x0Expected := 1.0
	x1Expected := 0.2

	TestBooleanValue("hasRoots", true, hasRoots, t)

	TestFloatValue("x0", x0Expected, x0, t)
	TestFloatValue("x1", x1Expected, x1, t)
}

func TestSolveQuadraticWithOneRoot(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(1, 2, 1)

	x0Expected := -1.0
	x1Expected := -1.0

	TestBooleanValue("hasRoots", true, hasRoots, t)

	TestFloatValue("x0", x0Expected, x0, t)
	TestFloatValue("x1", x1Expected, x1, t)
}

func TestSolveQuadraticWithOutRoots(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(1, -3, 4)

	x0Expected := 0.0
	x1Expected := 0.0

	TestBooleanValue("hasRoots", false, hasRoots, t)

	TestFloatValue("x0", x0Expected, x0, t)
	TestFloatValue("x1", x1Expected, x1, t)
}
