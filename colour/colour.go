package colour

import "github.com/rwbailey/ray/helpers"

type Colour struct {
	Red   float64
	Green float64
	Blue  float64
}

var (
	White, Black Colour
)

func init() {
	White = New(1, 1, 1)
	Black = New(0, 0, 0)
}

func New(r, g, b float64) Colour {
	return Colour{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}

func (c1 Colour) Add(c2 Colour) Colour {
	return Colour{
		Red:   c1.Red + c2.Red,
		Green: c1.Green + c2.Green,
		Blue:  c1.Blue + c2.Blue,
	}
}

func (c1 Colour) Subtract(c2 Colour) Colour {
	return Colour{
		Red:   c1.Red - c2.Red,
		Green: c1.Green - c2.Green,
		Blue:  c1.Blue - c2.Blue,
	}
}

func (c Colour) Multiply(i float64) Colour {
	return Colour{
		Red:   c.Red * i,
		Green: c.Green * i,
		Blue:  c.Blue * i,
	}
}

func (c1 Colour) Product(c2 Colour) Colour {
	return Colour{
		Red:   c1.Red * c2.Red,
		Green: c1.Green * c2.Green,
		Blue:  c1.Blue * c2.Blue,
	}
}

// Equals returns true if t1 == t2, else returns false
func (a Colour) Equals(b Colour) bool {
	return helpers.FloatEquals(a.Red, b.Red) && helpers.FloatEquals(a.Green, b.Green) && helpers.FloatEquals(a.Blue, b.Blue)
}
