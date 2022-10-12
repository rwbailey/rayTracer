package shape

import (
	"math"

	"github.com/rwbailey/ray/helpers"
	"github.com/rwbailey/ray/material"
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/tuple"
)

type Plane struct {
	Transform matrix.Matrix
	Material  *material.Material
}

func NewPlane() *Plane {
	return &Plane{
		Transform: matrix.IdentityMatrix(4),
		Material:  material.NewMaterial(),
	}
}

func (p *Plane) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	itm := p.Transform.Inverse()

	objNormal := tuple.Vector(0, 1, 0)

	worldNormal := itm.Transpose().Transform(objNormal)

	worldNormal.W = 0.0

	return worldNormal.Normalise()
}

func (p *Plane) Intersect(r ray.Ray) []*Intersection {

	inverseTransform := p.Transform.Inverse()
	rt := r.Transform(inverseTransform)

	if math.Abs(rt.Direction.Y) < helpers.Epsilon {
		return nil
	}

	t := -rt.Origin.Y / rt.Direction.Y
	return []*Intersection{
		{
			T:      t,
			Object: p,
		},
	}
}

func (p *Plane) SetMaterial(m *material.Material) {
	p.Material = m
}

func (p *Plane) GetMaterial() *material.Material {
	return p.Material
}

func (p *Plane) SetTransform(m matrix.Matrix) {
	p.Transform = m
}

func (p *Plane) GetTransform() matrix.Matrix {
	return p.Transform
}
