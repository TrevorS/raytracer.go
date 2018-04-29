package main

// Options represents rendering configuration.
type Options struct {
	width         int64
	height        int64
	fov           float64
	cameraToWorld Matrix44f
}
