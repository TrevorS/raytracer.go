package main

// Matrix44f represents a 4x4 matrix of floats.
type Matrix44f struct {
	numbers [4][4]float64
}

// Matrix44fNew returns a new matrix initialized to the passed in values.
func Matrix44fNew(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p float64) Matrix44f {
	x := [4][4]float64{
		[4]float64{a, b, c, d},
		[4]float64{e, f, g, h},
		[4]float64{i, j, k, l},
		[4]float64{m, n, o, p},
	}

	return Matrix44f{x}
}

// Matrix44fZero returns a matrix initialized to all zeros.
func Matrix44fZero() Matrix44f {
	return Matrix44fNew(
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	)
}

// Matrix44fIdentity returns the identity matrix.
func Matrix44fIdentity() Matrix44f {
	return Matrix44fNew(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)
}

func (m *Matrix44f) set(x, y int, value float64) {
	m.numbers[x][y] = value
}

func (m Matrix44f) get(x, y int) float64 {
	return m.numbers[x][y]
}

func (m Matrix44f) multiply(matrix Matrix44f) Matrix44f {
	result := Matrix44fZero()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value :=
				m.numbers[i][0]*matrix.numbers[0][j] +
					m.numbers[i][1]*matrix.numbers[1][j] +
					m.numbers[i][2]*matrix.numbers[2][j] +
					m.numbers[i][3]*matrix.numbers[3][j]

			result.numbers[i][j] = value
		}
	}

	return result
}

func (m Matrix44f) transpose() Matrix44f {
	result := Matrix44fZero()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			value := m.numbers[j][i]

			result.numbers[i][j] = value
		}
	}

	return result
}

func (m Matrix44f) multiplyVector(v Vec3f) Vec3f {
	x := v.x*m.numbers[0][0] + v.y*m.numbers[1][0] + v.z*m.numbers[2][0] + m.numbers[3][0]
	y := v.x*m.numbers[0][1] + v.y*m.numbers[1][1] + v.z*m.numbers[2][1] + m.numbers[3][1]
	z := v.x*m.numbers[0][2] + v.y*m.numbers[1][2] + v.z*m.numbers[2][2] + m.numbers[3][2]
	w := v.x*m.numbers[0][3] + v.y*m.numbers[1][3] + v.z*m.numbers[2][3] + m.numbers[3][3]

	if w != 1 && w != 0 {
		return Vec3f{x / w, y / w, z / w}
	}

	return Vec3f{x, y, z}
}

func (m Matrix44f) multiplyDirection(v Vec3f) Vec3f {
	x := v.x*m.numbers[0][0] + v.y*m.numbers[1][0] + v.z*m.numbers[2][0]
	y := v.x*m.numbers[0][1] + v.y*m.numbers[1][1] + v.z*m.numbers[2][1]
	z := v.x*m.numbers[0][2] + v.y*m.numbers[1][2] + v.z*m.numbers[2][2]

	return Vec3f{x, y, z}
}

func (m *Matrix44f) inverse() Matrix44f {
	result := Matrix44fIdentity()

	for i := 0; i < 3; i++ {
		pivot := i

		pivotSize := m.numbers[i][i]

		if pivotSize < 0 {
			pivotSize = -pivotSize

			for j := i + 1; j < 4; j++ {
				tmp := m.numbers[j][i]

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
			return Matrix44fIdentity()
		}

		if pivot != i {
			for j := 0; j < 4; j++ {
				m.numbers[i][j], m.numbers[pivot][j] = m.numbers[pivot][j], m.numbers[i][j]
				result.numbers[i][j], result.numbers[pivot][j] = result.numbers[pivot][j], result.numbers[i][j]
			}
		}

		for j := i + 1; j < 4; j++ {
			f := m.numbers[j][i] / m.numbers[i][i]

			for k := 0; k < 4; k++ {
				m.numbers[j][k] -= f * m.numbers[i][k]
				result.numbers[j][k] -= f * result.numbers[i][k]
			}
		}
	}

	for i := 3; i >= 0; i-- {
		f := m.numbers[i][i]

		if f == 0 {
			return Matrix44fIdentity()
		}

		for j := 0; j < 4; j++ {
			m.numbers[i][j] /= f
			result.numbers[i][j] /= f
		}

		for j := 0; j < i; j++ {
			f = m.numbers[j][i]

			for k := 0; k < 4; k++ {
				m.numbers[j][k] -= f * m.numbers[i][k]
				result.numbers[j][k] -= f * result.numbers[i][k]
			}
		}
	}

	return result
}
