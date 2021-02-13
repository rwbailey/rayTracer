package matrix

import (
	"math"
)

const epsilon = 0.00001

type Matrix [][]float64

func New(s [][]float64) Matrix {
	return Matrix(s)
}

func Zero(n int) Matrix {
	s := [][]float64{}
	for i := 0; i < n; i++ {
		s = append(s, make([]float64, n))
	}
	return New(s)
}

func Identity(n int) Matrix {
	z := Zero(n)
	for i := 0; i < n; i++ {
		z[i][i] = 1
	}
	return z
}

func (m1 Matrix) Equals(m2 Matrix) bool {
	for row := 0; row < len(m1); row++ {
		for col := 0; col < len(m1[row]); col++ {
			if !floatEquals(m1[row][col], m2[row][col]) {
				return false
			}
		}
	}
	return true
}

// Compare the equivelance of two floating point numbers to within the error margin epsilon
func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

// We only need to multiply 4x4 matrices
func (a Matrix) Multiply(b Matrix) Matrix {
	m := Zero(4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			m[row][col] = a[row][0]*b[0][col] +
				a[row][1]*b[1][col] +
				a[row][2]*b[2][col] +
				a[row][3]*b[3][col]
		}
	}

	return m
}
