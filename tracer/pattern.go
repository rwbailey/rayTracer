package tracer

import (
	"math"
)

type Pattern interface {
	// isPattern()
	ColourAt(p Tuple) Colour
	ColourAtObject(obj Shape, worldPoint Tuple) Colour
	SetTransform(m Matrix)
	GetTransform() Matrix
}

type Stripes struct {
	A         Colour
	B         Colour
	Transform Matrix
}

// func (*Stripes) isPattern() {}

func NewStripePattern(a, b Colour) *Stripes {
	return &Stripes{
		A:         a,
		B:         b,
		Transform: IdentityMatrix(4),
	}
}

func (s *Stripes) ColourAt(p Tuple) Colour {
	if int(math.Floor(p.X))%2 == 0 {
		return s.A
	}
	return s.B
}

func (s *Stripes) ColourAtObject(obj Shape, worldPoint Tuple) Colour {
	objPoint := obj.GetTransform().Inverse().MultiplyTuple(worldPoint)
	patternPoint := obj.GetMaterial().Pattern.GetTransform().Inverse().MultiplyTuple(objPoint)
	return obj.GetMaterial().Pattern.ColourAt(patternPoint)
}

func (s *Stripes) SetTransform(m Matrix) {
	s.Transform = m
}

func (s *Stripes) GetTransform() Matrix {
	return s.Transform
}
