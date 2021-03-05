package intersection

import (
	"sort"

	"github.com/rwbailey/ray/shape"
)

type Intersection struct {
	T      float64
	Object shape.Shape
}

func Intersections(ints ...*Intersection) []*Intersection {
	sort.Slice(ints, func(i, j int) bool { return ints[i].T < ints[j].T })
	return ints
}

// Assumes xs is sorted by T in assending order
func Hit(xs []*Intersection) *Intersection {
	for _, x := range xs {
		if x.T >= 0 {
			return x
		}
	}
	return nil
}
