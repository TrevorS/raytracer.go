package main

import (
	"math"
)

// Vec3f represents a three dimensional floating point vector.
type Vec3f struct {
	x, y, z float64
}

// Vec3fZero returns a Vector initialized to zero.
func Vec3fZero() Vec3f {
	return Vec3f{x: 0.0, y: 0.0, z: 0.0}
}

// Vec3fAt returns a Vector initialized with x, y, z as the same value.
func Vec3fAt(value float64) Vec3f {
	return Vec3f{x: value, y: value, z: value}
}

func (v Vec3f) negate() Vec3f {
	return Vec3f{
		x: -v.x,
		y: -v.y,
		z: -v.z,
	}
}

func (v Vec3f) add(vector Vec3f) Vec3f {
	return Vec3f{
		x: v.x + vector.x,
		y: v.y + vector.y,
		z: v.z + vector.z,
	}
}

func (v Vec3f) subtract(vector Vec3f) Vec3f {
	return Vec3f{
		x: v.x - vector.x,
		y: v.y - vector.y,
		z: v.z - vector.z,
	}
}

func (v Vec3f) multiply(vector Vec3f) Vec3f {
	return Vec3f{
		x: v.x * vector.x,
		y: v.y * vector.y,
		z: v.z * vector.z,
	}
}

func (v Vec3f) multiplyScalar(value float64) Vec3f {
	return Vec3f{
		x: value * v.x,
		y: value * v.y,
		z: value * v.z,
	}
}

func (v Vec3f) dotProduct(vector Vec3f) float64 {
	return v.x*vector.x + v.y*vector.y + v.z*vector.z
}

func (v Vec3f) divideScalar(value float64) Vec3f {
	return Vec3f{
		x: v.x / value,
		y: v.y / value,
		z: v.z / value,
	}
}

func (v Vec3f) crossProduct(vector Vec3f) Vec3f {
	return Vec3f{
		x: v.y*vector.z - v.z*vector.y,
		y: v.z*vector.x - v.x*vector.z,
		z: v.x*vector.y - v.y*vector.x,
	}
}

func (v Vec3f) norm() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3f) length() float64 {
	return math.Sqrt(v.norm())
}

func (v Vec3f) normalize() Vec3f {
	n := v.norm()

	if n > 0 {
		factor := 1 / math.Sqrt(n)

		return Vec3f{
			x: v.x * factor,
			y: v.y * factor,
			z: v.z * factor,
		}
	}

	return v
}
