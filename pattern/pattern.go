package pattern

import (
	"math"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/tuple"
)

type Pattern interface {
	// isPattern()
	ColourAt(p tuple.Tuple) colour.Colour
}

type Stripes struct {
	A colour.Colour
	B colour.Colour
}

// func (*Stripes) isPattern() {}

func NewStripePattern(a, b colour.Colour) *Stripes {
	return &Stripes{
		A: a,
		B: b,
	}
}

func (s *Stripes) ColourAt(p tuple.Tuple) colour.Colour {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}
	return s.B
}
