package tracer

import (
	"math"
)

type Sphere struct {
	Transform Matrix
	Material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: IdentityMatrix(4),
		Material:  NewMaterial(),
	}
}

func (s *Sphere) SetMaterial(m *Material) {
	s.Material = m
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}

func (s *Sphere) Intersect(r Ray) []*Intersection {

	inverseTransform := s.Transform.Inverse()

	rt := r.Transform(inverseTransform)

	var t1 float64
	var t2 float64

	sphereToRay := rt.Origin.Subtract(Point(0, 0, 0))
	a := rt.Direction.Dot(rt.Direction)
	b := 2 * rt.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - (4 * a * c)
	if discriminant < 0 {
		return nil
	}
	t1 = (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 = (-b + math.Sqrt(discriminant)) / (2 * a)

	return []*Intersection{
		{
			T:      t1,
			Object: s,
		},
		{
			T:      t2,
			Object: s,
		},
	}
}

func (s *Sphere) SetTransform(m Matrix) {
	s.Transform = m
}

func (s *Sphere) GetTransform() Matrix {
	return s.Transform
}

func (s *Sphere) NormalAt(worldPoint Tuple) Tuple {
	itm := s.Transform.Inverse()
	objPoint := itm.Transform(worldPoint)

	objNormal := objPoint.Subtract(Point(0, 0, 0))

	worldNormal := itm.Transpose().Transform(objNormal)

	worldNormal.W = 0.0

	return worldNormal.Normalise()
}
