package main

import (
	"testing"
)

func TestMatrix44fNew(t *testing.T) {
	a := 1.0
	b := 2.0
	c := 3.0
	d := 4.0
	e := 5.0
	f := 6.0
	g := 7.0
	h := 8.0
	i := 9.0
	j := 10.0
	k := 11.0
	l := 12.0
	m := 13.0
	n := 14.0
	o := 15.0
	p := 16.0

	matrix := Matrix44fNew(
		a, b, c, d,
		e, f, g, h,
		i, j, k, l,
		m, n, o, p,
	)

	TestMatrix44fValues(
		a, b, c, d,
		e, f, g, h,
		i, j, k, l,
		m, n, o, p,
		matrix,
		t,
	)
}

func TestMatrix44fZero(t *testing.T) {
	matrix := Matrix44fZero()

	TestMatrix44fValues(
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		matrix,
		t,
	)
}

func TestMatrix44fIdentity(t *testing.T) {
	matrix := Matrix44fIdentity()

	TestMatrix44fValues(
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
		matrix,
		t,
	)
}

func TestMatrix44fSet(t *testing.T) {
	matrix := Matrix44fZero()

	matrix.set(1, 2, 5.0)

	TestMatrix44fValues(
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 5.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		0.0, 0.0, 0.0, 0.0,
		matrix,
		t,
	)
}

func TestMatrix44fGet(t *testing.T) {
	matrix := Matrix44fZero()

	i := 2
	j := 3
	expected := 6.0

	matrix.set(i, j, expected)

	value := matrix.get(i, j)

	TestFloatValue("matrix set value", expected, value, t)
}

func TestMatrix44fMultiply(t *testing.T) {
	matrix1 := Matrix44fNew(
		1.0, 2.0, 3.0, 4.0,
		1.0, 2.0, 3.0, 4.0,
		1.0, 2.0, 3.0, 4.0,
		1.0, 2.0, 3.0, 4.0,
	)

	matrix2 := Matrix44fNew(
		4.0, 3.0, 2.0, 1.0,
		4.0, 3.0, 2.0, 1.0,
		4.0, 3.0, 2.0, 1.0,
		4.0, 3.0, 2.0, 1.0,
	)

	result := matrix1.multiply(matrix2)

	TestMatrix44fValues(
		40, 30, 20, 10,
		40, 30, 20, 10,
		40, 30, 20, 10,
		40, 30, 20, 10,
		result,
		t,
	)
}

func TestMatrix44fTranspose(t *testing.T) {
	matrix := Matrix44fNew(
		1.0, 2.0, 3.0, 4.0,
		1.0, 2.0, 3.0, 4.0,
		1.0, 2.0, 3.0, 4.0,
		1.0, 2.0, 3.0, 4.0,
	)

	result := matrix.transpose()

	TestMatrix44fValues(
		1.0, 1.0, 1.0, 1.0,
		2.0, 2.0, 2.0, 2.0,
		3.0, 3.0, 3.0, 3.0,
		4.0, 4.0, 4.0, 4.0,
		result,
		t,
	)
}

func TestMatrix44fInverse(t *testing.T) {
	matrix := Matrix44fNew(
		0.707107, 0, -0.707107, 0,
		-0.331295, 0.883452, -0.331295, 0,
		0.624695, 0.468521, 0.624695, 0,
		4.000574, 3.00043, 4.000574, 1,
	)

	result := matrix.inverse()

	TestMatrix44fValues(
		0.707107, -0.331295, 0.624695, 0,
		0, 0.883452, 0.468521, 0,
		-0.707107, -0.331295, 0.624695, 0,
		0, 0, -6.404044, 1,
		result,
		t,
	)
}

func TestMatrix44fInverseSingular(t *testing.T) {
	matrix := Matrix44fZero()

	result := matrix.inverse()

	TestMatrix44fValues(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
		result,
		t,
	)
}
