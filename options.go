package main

// Options represents rendering configuration.
type Options struct {
	width         int
	height        int
	fov           float64
	cameraToWorld Matrix44f
}
