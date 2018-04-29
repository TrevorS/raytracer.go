package main

import (
	"math"
)

// Clamp returns value restricted between the high and low parameters.
func Clamp(low, high, value float64) float64 {
	return math.Max(low, math.Min(high, value))
}

// DegreesToRadians converts degrees into radians.
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Mix mixes two color vectors and a value together and returns a new Vector.
func Mix(v1, v2 Vector, mixValue float64) Vector {
	return v1.multiplyScalar(1 - mixValue).add(v2.multiplyScalar(mixValue))
}
