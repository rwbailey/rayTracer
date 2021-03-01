package shape

import (
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/tuple"
)

type Shape interface {
	Intersect(r ray.Ray) []*Intersection
	SetTransform(m matrix.Matrix)
	NormalAt(worldPoint tuple.Tuple) tuple.Tuple
}
