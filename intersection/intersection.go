package intersection

import (
	"math"

	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
)

type Intersection struct {
	T      float64
	Object shape.Shape
}

func Intersect(s shape.Shape, r ray.Ray) []float64 {
	var t1 float64
	var t2 float64

	switch s.(type) {
	case *shape.Sphere:
		sphereToRay := r.Origin.Subtract(tuple.Point(0, 0, 0))
		a := r.Direction.Dot(r.Direction)
		b := 2 * r.Direction.Dot(sphereToRay)
		c := sphereToRay.Dot(sphereToRay) - 1

		discriminant := b*b - (4 * a * c)
		if discriminant < 0 {
			return nil
		}
		t1 = (-b - math.Sqrt(discriminant)) / (2 * a)
		t2 = (-b + math.Sqrt(discriminant)) / (2 * a)
	}
	return []float64{t1, t2}
}

func Intersections(ints ...Intersection) []Intersection {
	return ints
}
