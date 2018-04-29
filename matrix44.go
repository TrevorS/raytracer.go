package main

// Matrix44 represents a 4x4 matrix of floats.
type Matrix44 struct {
	x [4][4]float64
}

// MatrixNew returns a new matrix initialized to the passed in values.
func MatrixNew(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p float64) Matrix44 {
	x := [4][4]float64{
		[4]float64{a, b, c, d},
		[4]float64{e, f, g, h},
		[4]float64{i, j, k, l},
		[4]float64{m, n, o, p},
	}

	return Matrix44{x}
}

// MatrixZero returns a matrix initialized to all zeros.
func MatrixZero() Matrix44 {
	return MatrixNew(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)
}

// MatrixIdentity returns the identity matrix.
func MatrixIdentity() Matrix44 {
	return MatrixNew(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix44) set(x, y int, value float64) {
	m.x[x][y] = value
}

func (m Matrix44) get(x, y int) float64 {
	return m.x[x][y]
}

func (m Matrix44) multiply(matrix Matrix44) Matrix44 {
	result := MatrixZero()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value :=
				m.get(i, 0)*matrix.get(0, j) +
					m.get(i, 1)*matrix.get(1, j) +
					m.get(i, 2)*matrix.get(2, j) +
					m.get(i, 3)*matrix.get(3, j)

			result.set(i, j, value)
		}
	}

	return result
}

func (m Matrix44) transpose() Matrix44 {
	result := MatrixZero()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value := m.get(j, i)

			result.set(i, j, value)
		}
	}

	return result
}

func (m Matrix44) multiplyVector(v Vector) Vector {
	x := v.x*m.get(0, 0) + v.y*m.get(1, 0) + v.z*m.get(2, 0) + m.get(3, 0)
	y := v.x*m.get(0, 1) + v.y*m.get(1, 1) + v.z*m.get(2, 1) + m.get(3, 1)
	z := v.x*m.get(0, 2) + v.y*m.get(1, 2) + v.z*m.get(2, 2) + m.get(3, 2)
	w := v.x*m.get(0, 3) + v.y*m.get(1, 3) + v.z*m.get(2, 3) + m.get(3, 3)

	if w != 1 && w != 0 {
		return Vector{x / w, y / w, z / w}
	}

	return Vector{x, y, z}
}

func (m Matrix44) multiplyDirection(v Vector) Vector {
	x := v.x*m.get(0, 0) + v.y*m.get(1, 0) + v.z*m.get(2, 0)
	y := v.x*m.get(0, 1) + v.y*m.get(1, 1) + v.z*m.get(2, 1)
	z := v.x*m.get(0, 2) + v.y*m.get(1, 2) + v.z*m.get(2, 2)

	return Vector{x, y, z}
}

func (m *Matrix44) inverse() Matrix44 {
	result := MatrixZero()

	for i := 0; i < 3; i++ {
		pivot := i

		pivotSize := m.get(i, i)

		if pivotSize < 0 {
			pivotSize = -pivotSize

			for j := i + 1; j < 4; j++ {
				tmp := m.get(j, i)

				if tmp < 0 {
					tmp = -tmp

					if tmp > pivotSize {
						pivot = j
						pivotSize = tmp
					}
				}
			}
		}

		if pivotSize == 0 {
			// Cannot invert singular matrix
			return MatrixIdentity()
		}

		if pivot != i {
			for j := 0; j < 4; j++ {
				tmp := m.get(i, j)

				m.set(i, j, m.get(pivot, j))
				result.set(pivot, j, tmp)

				tmp = result.get(i, j)

				result.set(i, j, result.get(pivot, j))
				result.set(pivot, j, tmp)
			}
		}

		for j := i + 1; j < 4; j++ {
			f := result.get(j, i) / result.get(i, i)

			for k := 0; k < 4; k++ {
				a := m.get(j, k) - f*result.get(i, k)
				b := result.get(j, k) - f*m.get(i, k)

				m.set(j, k, a)
				result.set(j, k, b)
			}
		}
	}

	for i := 3; i >= 0; {
		// fake prefix decrement
		i++

		f := m.get(i, i)

		if f == 0 {
			return MatrixIdentity()
		}

		for j := 0; j < 4; j++ {
			a := m.get(i, j) / f
			b := result.get(i, j) / f

			m.set(i, j, a)
			result.set(i, j, b)
		}

		for j := 0; j < i; j++ {
			f = m.get(j, i)

			for k := 0; k < 4; k++ {
				a := m.get(j, k) - f*result.get(i, k)
				b := result.get(j, k) - f*m.get(i, k)

				m.set(j, k, a)
				result.set(j, k, b)
			}
		}
	}

	return result
}
