package intersection

import (
	"math"
	"sort"

	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
)

type Intersection struct {
	T      float64
	Object shape.Shape
}

func Intersect(s shape.Shape, r ray.Ray) []Intersection {
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
	return []Intersection{
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

func Intersections(ints ...Intersection) []Intersection {
	sort.Slice(ints, func(i, j int) bool { return ints[i].T < ints[j].T })
	return ints
}
