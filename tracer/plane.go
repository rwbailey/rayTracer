package tracer

import (
	"math"
)

type Plane struct {
	Transform Matrix
	Material  *Material
}

func NewPlane() *Plane {
	return &Plane{
		Transform: IdentityMatrix(4),
		Material:  NewMaterial(),
	}
}

func (p *Plane) NormalAt(worldPoint Tuple) Tuple {
	itm := p.Transform.Inverse()

	objNormal := Vector(0, 1, 0)

	worldNormal := itm.Transpose().Transform(objNormal)

	worldNormal.W = 0.0

	return worldNormal.Normalise()
}

func (p *Plane) Intersect(r Ray) []*Intersection {

	inverseTransform := p.Transform.Inverse()
	rt := r.Transform(inverseTransform)

	if math.Abs(rt.Direction.Y) < Epsilon {
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

func (p *Plane) SetMaterial(m *Material) {
	p.Material = m
}

func (p *Plane) GetMaterial() *Material {
	return p.Material
}

func (p *Plane) SetTransform(m Matrix) {
	p.Transform = m
}

func (p *Plane) GetTransform() Matrix {
	return p.Transform
}
