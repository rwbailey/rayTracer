package shape

import (
	"github.com/rwbailey/ray/material"
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/tuple"
)

type Shape interface {
	GetMaterial() *material.Material
	GetTransform() matrix.Matrix
	Intersect(r ray.Ray) []*Intersection
	NormalAt(worldPoint tuple.Tuple) tuple.Tuple
	SetTransform(m matrix.Matrix)
}
