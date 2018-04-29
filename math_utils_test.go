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
	v1 := Vector{1, 2, 3}
	v2 := Vector{2, 3, 4}
	mixValue := 2.0

	result := Mix(v1, v2, mixValue)

	expected := Vector{3, 4, 5}

	if result != expected {
		t.Errorf("%v should have been %v", result, expected)
	}
}
