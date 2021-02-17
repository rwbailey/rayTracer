package ray

import (
	"github.com/rwbailey/ray/tuple"
)

type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

func New(origin, direction tuple.Tuple) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) Position(t float64) tuple.Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}
