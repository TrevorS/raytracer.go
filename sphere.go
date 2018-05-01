package main

import (
	"math"
	"math/rand"
)

// Sphere is a GeometricObject with a center and a radius.
type Sphere struct {
	center  Vec3f
	color   Vec3f
	radius  float64
	radius2 float64
}

// NewSphere returns a Sphere with the radius squared pre-computed.
// You should always use this to make new Spheres.
func NewSphere(center Vec3f, radius float64) Sphere {
	color := generateRandomColor()

	return Sphere{
		center:  center,
		color:   color,
		radius:  radius,
		radius2: radius * radius,
	}
}

func (s Sphere) intersect(origin, direction Vec3f) (intersects bool, t float64) {
	L := s.center.subtract(origin)

	tca := L.dotProduct(direction)

	if tca < 0 {
		return false, math.Inf(1)
	}

	d2 := L.dotProduct(L) - tca*tca

	if d2 > s.radius2 {
		return false, math.Inf(1)
	}

	thc := math.Sqrt(s.radius2 - d2)

	t0 := tca - thc
	t1 := tca + thc

	if t0 > t1 {
		t1, t0 = t0, t1
	}

	if t0 < 0 {
		t0 = t1

		if t0 < 0 {
			return false, math.Inf(1)
		}
	}

	return true, t0
}

func (s Sphere) getSurfaceData(phit Vec3f) (nhit Vec3f, tex Vec2f) {
	nhit = phit.subtract(s.center)
	nhit.normalize()

	x := (1 + math.Atan2(nhit.z, nhit.x)/math.Pi) * 0.5
	y := math.Acos(nhit.y) / math.Pi

	return nhit, Vec2f{x: x, y: y}
}

func (s Sphere) getColor() Vec3f {
	return s.color
}

func generateRandomColor() Vec3f {
	return Vec3f{
		x: rand.Float64(),
		y: rand.Float64(),
		z: rand.Float64(),
	}
}
