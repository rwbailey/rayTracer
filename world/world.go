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

func NewWorld() *World {
	return &World{
		Objects: []shape.Shape{},
		Light:   nil,
	}
}

func DefaultWorld() *World {
	ls := light.NewPointLight(tuple.Point(-10, 10, -10), colour.NewColour(1, 1, 1))

	s1 := shape.NewSphere()
	s1.Material.Colour = colour.NewColour(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := shape.NewSphere()
	s2.Transform = matrix.ScalingMatrix(0.5, 0.5, 0.5)

	w := NewWorld()
	w.Light = ls
	w.Objects = []shape.Shape{
		s1,
		s2,
	}

	return w
}

func (w *World) AddObjects(objects ...shape.Shape) {
	w.Objects = append(w.Objects, objects...)
}

// IntersectWorld iterates over all of the objects in the world, and returns
// a sorted slice of the Intersections
func (w *World) IntersectWorld(r ray.Ray) []*shape.Intersection {
	xs := make([]*shape.Intersection, 0, len(w.Objects))
	for _, s := range w.Objects {
		xs = append(xs, s.Intersect(r)...)
	}
	return shape.Intersections(xs...)
}

// Given a set of pre computed values for an intersection, calculate the lighting/colour
// at that point in the world.
func (w *World) ShadeHit(c *shape.Computations) colour.Colour {
	return c.Object.GetMaterial().Lighting(
		w.Light,
		c.Point,
		c.Eyev,
		c.Normalv,
		w.IsShadowed(c.OverPoint),
	)
}

func (w *World) ColourAt(r ray.Ray) colour.Colour {
	xs := w.IntersectWorld(r)
	hit := shape.Hit(xs)
	if hit == nil {
		return colour.NewColour(0, 0, 0)
	}
	c := hit.PrepareComputations(r)
	return w.ShadeHit(c)
}

func (w *World) IsShadowed(p tuple.Tuple) bool {
	v := w.Light.Position.Subtract(p)

	distance := v.Magnitude()
	direction := v.Normalise()
	r := ray.NewRay(p, direction)
	xs := w.IntersectWorld(r)
	h := shape.Hit(xs)
	if h != nil && h.T < distance {
		return true
	}
	return false
}
