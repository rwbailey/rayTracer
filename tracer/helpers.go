package tracer

import (
	"math"
)

const Epsilon = 0.00001

// Compare the equivelance of two floating point numbers to to the given precision epsilon
func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}
