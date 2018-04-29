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

// SolveQuadratic solves the quadratic equation.
// Returns:
// hasRoots: true if roots / false if not
// x0, x1: the roots / 0 if no roots
func SolveQuadratic(a, b, c float64) (hasRoots bool, x0, x1 float64) {
	hasRoots, x0, x1 = false, 0, 0
	discr := b*b - 4*a*c

	if discr < 0 {
		return
	} else if discr == 0 {
		x0 = -0.5 * b / a
		x1 = x0
	} else {
		var q float64

		if b > 0 {
			q = -0.5 * (b + math.Sqrt(discr))
		} else {
			q = -0.5 * (b - math.Sqrt(discr))
		}

		x0 = q / a
		x1 = c / q
	}
	hasRoots = true

	return
}
