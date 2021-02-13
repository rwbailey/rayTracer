package matrix

import "math"

const epsilon = 0.00001

type Matrix [][]float64

func New(s [][]float64) Matrix {
	return Matrix(s)
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
