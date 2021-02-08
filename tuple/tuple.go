package tuple

import "math"

const epsilon = 0.00001

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// New returns a new tuple
func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

// Point returns a new point
func Point(x, y, z float64) Tuple {
	return New(x, y, z, 1.0)
}

// Vector returns new vector
func Vector(x, y, z float64) Tuple {
	return New(x, y, z, 0.0)
}

// IsPoint returns true if t is a point, else returns false
func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector returns true if t is a vector, else returns false
func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

// Equals returns true if t1 == t2, else returns false
func (t1 Tuple) Equals(t2 Tuple) bool {
	return floatEquals(t1.X, t2.X) && floatEquals(t1.Y, t2.Y) && floatEquals(t1.Z, t2.Z) && floatEquals(t1.W, t2.W)
}

// Add returns t1 + t2
func (t1 Tuple) Add(t2 Tuple) Tuple {
	return Tuple{
		X: t1.X + t2.X,
		Y: t1.Y + t2.Y,
		Z: t1.Z + t2.Z,
		W: t1.W + t2.W,
	}
}

// Subtract returns t1 - t2
func (t1 Tuple) Subtract(t2 Tuple) Tuple {
	return Tuple{
		X: t1.X - t2.X,
		Y: t1.Y - t2.Y,
		Z: t1.Z - t2.Z,
		W: t1.W - t2.W,
	}
}

// Negate returns -t when given t
func (t Tuple) Negate() Tuple {
	return Tuple{
		X: t.X * -1,
		Y: t.Y * -1,
		Z: t.Z * -1,
		W: t.W * -1, // Do we need this?
	}
}

// Compare the equivelance of two floating point numbers to within the error margin epsilon
func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}
