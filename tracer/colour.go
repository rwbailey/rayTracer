package tracer

type Colour struct {
	Red   float64
	Green float64
	Blue  float64
}

var (
	White, Black, Red, Green, Blue, Yellow, Brown, Gold, Cyan, Grey, Orange, Pink, Purple Colour
)

func init() {
	White = NewColour(1, 1, 1)
	Black = NewColour(0, 0, 0)
	Red = NewColour(1, 0, 0)
	Green = NewColour(0, 1, 0)
	Blue = NewColour(0, 0, 1)
	Yellow = NewColour(1, 1, 0)
	Brown = NewColour(165.0/255, 42.0/255, 42.0/255)
	Gold = NewColour(1, 215.0/255, 0)
	Cyan = NewColour(0, 1, 1)
	Grey = NewColour(128.0/255, 128.0/255, 128.0/255)
	Orange = NewColour(1, 140.0/255, 0)
	Pink = NewColour(1, 192.0/255, 203.0/255)
	Purple = NewColour(128.0/255, 0, 128.0/255)
}

func NewColour(r, g, b float64) Colour {
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

func (a Colour) Equals(b Colour) bool {
	return FloatEquals(a.Red, b.Red) && FloatEquals(a.Green, b.Green) && FloatEquals(a.Blue, b.Blue)
}
