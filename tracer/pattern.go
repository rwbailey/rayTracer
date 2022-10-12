package tracer

import (
	"math"
)

type Pattern interface {
	isPattern()
	ColourAt(p Tuple) Colour
}

type Stripes struct {
	A Colour
	B Colour
}

func (*Stripes) isPattern() {}

func NewStripePattern(a, b Colour) *Stripes {
	return &Stripes{
		A: a,
		B: b,
	}
}

func (s *Stripes) ColourAt(p Tuple) Colour {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}
	return s.B
}
