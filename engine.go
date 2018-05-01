package main

import (
	"math"
)

// Render renders a scene of GeometricObjects according to a set of Options.
func Render(options Options, objects []GeometricObject) []Vec3f {
	framebuffer := make([]Vec3f, options.width*options.height)

	fHeight := float64(options.height)
	fWidth := float64(options.width)

	scale := math.Tan(DegreesToRadians(options.fov * 0.5))
	aspectRatio := fWidth / fHeight

	origin := options.cameraToWorld.multiplyVector(Vec3fZero())

	for j := 0; j < options.height; j++ {
		fJ := float64(j)
		for i := 0; i < options.width; i++ {
			fI := float64(i)

			x := (2*(fI+0.5)/fWidth - 1) * aspectRatio * scale
			y := (1 - 2*(fJ+0.5)) / fHeight * scale

			direction := options.cameraToWorld.multiplyDirection(Vec3f{x, y, -1.0})
			direction.normalize()

			ray := castRay(origin, direction, objects)

			framebuffer = append(framebuffer, ray)
		}
	}

	return framebuffer
}

func castRay(origin, direction Vec3f, objects []GeometricObject) Vec3f {
	hitColor := Vec3fZero()

	didHit, t, hitObject := trace(origin, direction, objects)

	if didHit {
		phit := origin.add(direction.multiplyScalar(t))

		nhit, tex := hitObject.getSurfaceData(phit)

		scale := 4.0

		pattern := getPattern(tex, scale)

		colorMix := Mix(hitObject.getColor(), hitObject.getColor().multiplyScalar(0.8), pattern)

		hitColor = colorMix.multiplyScalar(math.Max(0.0, nhit.dotProduct(direction.multiplyScalar(-1))))
	}

	return hitColor
}

func trace(origin, direction Vec3f, objects []GeometricObject) (ditHit bool, tNear float64, hitObject GeometricObject) {
	tNear = math.Inf(1)

	for _, object := range objects {
		intersects, t := object.intersect(origin, direction)

		if intersects && t < tNear {
			hitObject = object
			tNear = t
		}
	}

	return hitObject != nil, tNear, hitObject
}

func getPattern(tex Vec2f, scale float64) float64 {
	var p1, p2 int

	if math.Mod(tex.x*scale, 1) > 0.5 {
		p1 = 1
	} else {
		p1 = 0
	}

	if math.Mod(tex.y*scale, 1) > 0.5 {
		p2 = 1
	} else {
		p2 = 0
	}

	return float64(p1 ^ p2)
}
