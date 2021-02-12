package colour

import "math"

const epsilon = 0.00001

type colour struct {
	Red   float64
	Green float64
	Blue  float64
}

func New(r, g, b float64) colour {
	return colour{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

func (c1 colour) Add(c2 colour) colour {
	return colour{
		Red:   c1.Red + c2.Red,
		Green: c1.Green + c2.Green,
		Blue:  c1.Blue + c2.Blue,
	}
}

func (c1 colour) Subtract(c2 colour) colour {
	return colour{
		Red:   c1.Red - c2.Red,
		Green: c1.Green - c2.Green,
		Blue:  c1.Blue - c2.Blue,
	}
}

func (c colour) Multiply(i float64) colour {
	return colour{
		Red:   c.Red * i,
		Green: c.Green * i,
		Blue:  c.Blue * i,
	}
}

func (c1 colour) Product(c2 colour) colour {
	return colour{
		Red:   c1.Red * c2.Red,
		Green: c1.Green * c2.Green,
		Blue:  c1.Blue * c2.Blue,
	}
}

// Equals returns true if t1 == t2, else returns false
func (a colour) Equals(b colour) bool {
	return floatEquals(a.Red, b.Red) && floatEquals(a.Green, b.Green) && floatEquals(a.Blue, b.Blue)
}

// Compare the equivelance of two floating point numbers to within the error margin epsilon
func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}
