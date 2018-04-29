package main

import (
	"testing"
)

func TestClampHigh(t *testing.T) {
	low := 2.0
	high := 4.0
	value := 6.0

	result := Clamp(low, high, value)

	if result != high {
		t.Errorf("%v should have been %v", value, high)
	}
}

func TestClampLow(t *testing.T) {
	low := 2.0
	high := 4.0
	value := 1.0

	result := Clamp(low, high, value)

	if result != low {
		t.Errorf("%v should have been %v", value, low)
	}
}

func TestDegreesToRadians(t *testing.T) {
	degrees := 45.0
	radians := 0.7853981633974483

	result := DegreesToRadians(degrees)

	if result != radians {
		t.Errorf("%v should have been %v", result, radians)
	}
}

func TestMix(t *testing.T) {
	v1 := Vec3f{1, 2, 3}
	v2 := Vec3f{2, 3, 4}
	mixValue := 2.0

	result := Mix(v1, v2, mixValue)

	expected := Vec3f{3, 4, 5}

	if result != expected {
		t.Errorf("%v should have been %v", result, expected)
	}
}

func TestSolveQuadraticWithTwoRoots(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(5, 6, 1)

	x0Expected := -1.0
	x1Expected := -0.2

	if hasRoots != true {
		t.Errorf("hasRoots should be %v", true)
	}

	if x0 != x0Expected {
		t.Errorf("%v should have been %v", x0, x0Expected)
	}

	if x1 != x1Expected {
		t.Errorf("%v should have been %v", x1, x1Expected)
	}
}

func TestSolveQuadraticWithTwoRootsNegativeB(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(5, -6, 1)

	x0Expected := 1.0
	x1Expected := 0.2

	if hasRoots != true {
		t.Errorf("hasRoots should be %v", true)
	}

	if x0 != x0Expected {
		t.Errorf("%v should have been %v", x0, x0Expected)
	}

	if x1 != x1Expected {
		t.Errorf("%v should have been %v", x1, x1Expected)
	}
}

func TestSolveQuadraticWithOneRoot(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(1, 2, 1)

	x0Expected := -1.0
	x1Expected := -1.0

	if hasRoots != true {
		t.Errorf("hasRoots should be %v", true)
	}

	if x0 != x0Expected {
		t.Errorf("%v should have been %v", x0, x0Expected)
	}

	if x1 != x1Expected {
		t.Errorf("%v should have been %v", x1, x1Expected)
	}
}

func TestSolveQuadraticWithOutRoots(t *testing.T) {
	hasRoots, x0, x1 := SolveQuadratic(1, -3, 4)

	x0Expected := 0.0
	x1Expected := 0.0

	if hasRoots != false {
		t.Errorf("hasRoots should be %v", false)
	}

	if x0 != x0Expected {
		t.Errorf("%v should have been %v", x0, x0Expected)
	}

	if x1 != x1Expected {
		t.Errorf("%v should have been %v", x1, x1Expected)
	}
}
