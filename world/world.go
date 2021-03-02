package world

import (
	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
)

type World struct {
	Objects []shape.Shape
	Light   *light.PointLight
}

func New() *World {
	return &World{
		Objects: []shape.Shape{},
		Light:   nil,
	}
}

func Default() *World {
	ls := light.NewPointLight(tuple.Point(-10, 10, -10), colour.New(1, 1, 1))

	s1 := shape.NewSphere()
	s1.Material.Colour = colour.New(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := shape.NewSphere()
	s2.Transform = matrix.Scaling(0.5, 0.5, 0.5)

	w := New()

	w.Light = ls

	w.Objects = []shape.Shape{
		s1,
		s2,
	}

	return w
}

func (w *World) IntersectWorld(r ray.Ray) []*shape.Intersection {
	xs := make([]*shape.Intersection, 0, len(w.Objects))
	for _, s := range w.Objects {
		xs = append(xs, s.Intersect(r)...)
	}

	return shape.Intersections(xs...)
}
