package tracer

import (
	"sort"
)

type Intersection struct {
	T      float64
	Object Shape
}

type Computations struct {
	T         float64
	Object    Shape
	Point     Tuple
	Eyev      Tuple
	Normalv   Tuple
	Inside    bool
	OverPoint Tuple
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

func (i *Intersection) PrepareComputations(r Ray) *Computations {
	c := Computations{}
	c.T, c.Object = i.T, i.Object

	c.Point = r.Position(c.T)
	c.Eyev = r.Direction.Negate()
	c.Normalv = c.Object.NormalAt(c.Point)

	if c.Normalv.Dot(c.Eyev) < 0 {
		c.Inside = true
		c.Normalv = c.Normalv.Negate()
	}
	c.OverPoint = c.Point.Add(c.Normalv.Multiply(Epsilon))

	return &c
}
