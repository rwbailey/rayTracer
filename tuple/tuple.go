package tuple

import "math"

const epsilon = 0.00001

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func New(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float64) Tuple {
	return New(x, y, z, 1.0)
}

func Vector(x, y, z float64) Tuple {
	return New(x, y, z, 0.0)
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t1 Tuple) Equals(t2 Tuple) bool {
	return equals(t1.X, t2.X) && equals(t1.Y, t2.Y) && equals(t1.Z, t2.Z) && equals(t1.W, t2.W)
}

func equals(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

func (t1 Tuple) Add(t2 Tuple) Tuple {
	return Tuple{
		X: t1.X + t2.X,
		Y: t1.Y + t2.Y,
		Z: t1.Z + t2.Z,
		W: t1.W + t2.W,
	}
}

func (t1 Tuple) Subtract(t2 Tuple) Tuple {
	return Tuple{
		X: t1.X - t2.X,
		Y: t1.Y - t2.Y,
		Z: t1.Z - t2.Z,
		W: t1.W - t2.W,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		X: t.X * -1,
		Y: t.Y * -1,
		Z: t.Z * -1,
		W: t.W * -1,
	}
}
