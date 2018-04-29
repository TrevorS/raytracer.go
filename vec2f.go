package main

// Vec2f represents a two dimensional floating point vector.
type Vec2f struct {
	x, y float64
}

// Vec2fZero returns a Vector initialized to zero.
func Vec2fZero() Vec2f {
	return Vec2f{x: 0, y: 0}
}

// Vec2fAt returns a Vector initialized with x, y as the same value.
func Vec2fAt(value float64) Vec2f {
	return Vec2f{x: 0, y: 0}
}

func (v Vec2f) add(vector Vec2f) Vec2f {
	return Vec2f{
		x: v.x + vector.x,
		y: v.y + vector.y,
	}
}

func (v Vec2f) subtract(vector Vec2f) Vec2f {
	return Vec2f{
		x: v.x - vector.x,
		y: v.y - vector.y,
	}
}

func (v Vec2f) multiply(vector Vec2f) Vec2f {
	return Vec2f{
		x: v.x * vector.x,
		y: v.y * vector.y,
	}
}

func (v Vec2f) divide(vector Vec2f) Vec2f {
	return Vec2f{
		x: v.x / vector.x,
		y: v.y / vector.y,
	}
}

func (v Vec2f) multiplyScalar(value float64) Vec2f {
	return Vec2f{
		x: value * v.x,
		y: value * v.y,
	}
}

func (v Vec2f) divideScalar(value float64) Vec2f {
	return Vec2f{
		x: v.x / value,
		y: v.y / value,
	}
}

func (v *Vec2f) multiplyInPlace(vector Vec2f) *Vec2f {
	v.x *= vector.x
	v.y *= vector.y

	return v
}

func (v *Vec2f) divideInPlace(vector Vec2f) *Vec2f {
	v.x /= vector.x
	v.y /= vector.y

	return v
}
