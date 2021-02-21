package shape

import (
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
)

type Shape interface {
	Intersect(r ray.Ray) []*Intersection
	SetTransform(m matrix.Matrix)
}
