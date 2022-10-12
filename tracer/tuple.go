package tracer

import (
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

// NewTuple returns a new tuple
func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

// Point returns a new point
func Point(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1.0)
}

// Vector returns new vector
func Vector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0.0)
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
	return FloatEquals(t1.X, t2.X) && FloatEquals(t1.Y, t2.Y) && FloatEquals(t1.Z, t2.Z) && FloatEquals(t1.W, t2.W)
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
		W: t.W * -1,
	}
}

func (t Tuple) Multiply(a float64) Tuple {

	return Tuple{
		X: t.X * a,
		Y: t.Y * a,
		Z: t.Z * a,
		W: t.W * a,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func (t Tuple) Normalise() Tuple {
	x := t.X / t.Magnitude()
	y := t.Y / t.Magnitude()
	z := t.Z / t.Magnitude()

	return Vector(x, y, z)
}

func (a Tuple) Dot(b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Tuple) Cross(b Tuple) Tuple {
	return Vector(
		a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X,
	)
}

func (v Tuple) Reflect(n Tuple) Tuple {
	return v.Subtract(n.Multiply(2 * v.Dot(n)))
}
