package helpers

import "math"

const epsilon = 0.00001

// Compare the equivelance of two floating point numbers to within the error margin epsilon
func FloatEquals(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}
