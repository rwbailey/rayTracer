package shape

import (
	"log"
	"math"

	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/tuple"
)

type Sphere struct {
	Transform matrix.Matrix
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: matrix.Identity(4),
	}
}

func (s *Sphere) Intersect(r ray.Ray) []*Intersection {

	inverseTransform, err := s.Transform.Inverse()
	if err != nil {
		log.Fatal(err)
	}

	rt := r.Transform(inverseTransform)

	var t1 float64
	var t2 float64

	sphereToRay := rt.Origin.Subtract(tuple.Point(0, 0, 0))
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

func (s *Sphere) SetTransform(m matrix.Matrix) {
	s.Transform = m
}

func (s *Sphere) NormalAt(worldPoint tuple.Tuple) tuple.Tuple {
	itm, _ := s.Transform.Inverse()
	objPoint := itm.Transform(worldPoint)

	objNormal := objPoint.Subtract(tuple.Point(0, 0, 0))

	worldNormal := itm.Transpose().Transform(objNormal)

	worldNormal.W = 0.0

	return worldNormal.Normalise()
}
