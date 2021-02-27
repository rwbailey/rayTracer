package light

import (
	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/tuple"
)

type PointLight struct {
	Position  tuple.Tuple
	Intensity colour.Colour
}

func NewPointLight(p tuple.Tuple, i colour.Colour) *PointLight {
	return &PointLight{
		Position:  p,
		Intensity: i,
	}
}
