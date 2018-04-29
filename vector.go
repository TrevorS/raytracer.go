package main

import (
	"math"
)

// Vector represents a three dimensional floating point vector.
type Vector struct {
	x, y, z float64
}

// VectorZero returns a Vector initialized to zero.
func VectorZero() Vector {
	return Vector{x: 0.0, y: 0.0, z: 0.0}
}

// VectorAt returns a Vector initialized with x, y, z as the same value.
func VectorAt(value float64) Vector {
	return Vector{x: value, y: value, z: value}
}

func (v Vector) negate() Vector {
	return Vector{
		x: -v.x,
		y: -v.y,
		z: -v.z,
	}
}

func (v Vector) add(vector Vector) Vector {
	return Vector{
		x: v.x + vector.x,
		y: v.y + vector.y,
		z: v.z + vector.z,
	}
}

func (v Vector) subtract(vector Vector) Vector {
	return Vector{
		x: v.x - vector.x,
		y: v.y - vector.y,
		z: v.z - vector.z,
	}
}

func (v Vector) multiply(vector Vector) Vector {
	return Vector{
		x: v.x * vector.x,
		y: v.y * vector.y,
		z: v.z * vector.z,
	}
}

func (v Vector) multiplyScalar(value float64) Vector {
	return Vector{
		x: value * v.x,
		y: value * v.y,
		z: value * v.z,
	}
}

func (v Vector) dotProduct(vector Vector) float64 {
	return v.x*vector.x + v.y*vector.y + v.z*vector.z
}

func (v Vector) divideScalar(value float64) Vector {
	return Vector{
		x: v.x / value,
		y: v.y / value,
		z: v.z / value,
	}
}

func (v Vector) crossProduct(vector Vector) Vector {
	return Vector{
		x: v.y*vector.z - v.z*vector.y,
		y: v.z*vector.x - v.x*vector.z,
		z: v.x*vector.y - v.y*vector.x,
	}
}

func (v Vector) norm() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vector) length() float64 {
	return math.Sqrt(v.norm())
}

func (v Vector) normalize() Vector {
	n := v.norm()

	if n > 0 {
		factor := 1 / math.Sqrt(n)

		return Vector{
			x: v.x * factor,
			y: v.y * factor,
			z: v.z * factor,
		}
	}

	return v
}
