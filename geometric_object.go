package main

// GeometricObject is an interface that requires implementation
// of intersection of surface logic.
type GeometricObject interface {
	intersect(origin, direction Vector) (bool, float64)
	getSurfaceData(phit Vector) (Vector, Vector)
}
