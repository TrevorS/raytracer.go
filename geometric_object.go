package main

// GeometricObject is an interface that requires implementation
// of intersection of surface logic.
type GeometricObject interface {
	intersect(origin, direction Vec3f) (bool, float64)
	getSurfaceData(phit Vec3f) (Vec3f, Vec3f)
}
