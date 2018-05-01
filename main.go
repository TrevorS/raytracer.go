package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("RayTracer.go")

	numberOfSpheres := 32

	objects := make([]GeometricObject, numberOfSpheres)

	for i := 0; i < numberOfSpheres; i++ {
		randomPosition := getRandomPosition()
		randomRadius := getRandomRadius()

		objects[i] = NewSphere(randomPosition, randomRadius)
	}

	cameraToWorld := getCameraToWorld()

	options := Options{
		width:         640,
		height:        480,
		fov:           51.52,
		cameraToWorld: cameraToWorld,
	}

	framebuffer := Render(options, objects)

	WriteImage("output.ppm", options, framebuffer)
}

func getRandomPosition() Vec3f {
	x := (0.5 - rand.Float64()*10)
	y := (0.5 - rand.Float64()*10)
	z := (0.5 + rand.Float64()*10)

	return Vec3f{x, y, z}
}

func getRandomRadius() float64 {
	return (0.5 + rand.Float64()*0.5)
}

func getCameraToWorld() Matrix44f {
	return Matrix44fNew(
		0.945519,
		0,
		-0.325569,
		0,
		-0.179534,
		0.834209,
		-0.521403,
		0,
		0.271593,
		0.551447,
		0.78876,
		0,
		4.208271,
		8.374532,
		17.932925,
		1,
	)
}
