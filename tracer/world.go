package tracer

type World struct {
	Objects []Shape
	Light   *PointLight
}

func NewWorld() *World {
	return &World{
		Objects: []Shape{},
		Light:   nil,
	}
}

func DefaultWorld() *World {
	ls := NewPointLight(Point(-10, 10, -10), NewColour(1, 1, 1))

	s1 := NewSphere()
	s1.Material.Colour = NewColour(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := NewSphere()
	s2.Transform = ScalingMatrix(0.5, 0.5, 0.5)

	w := NewWorld()
	w.Light = ls
	w.Objects = []Shape{
		s1,
		s2,
	}

	return w
}

func (w *World) AddObjects(objects ...Shape) {
	w.Objects = append(w.Objects, objects...)
}

// IntersectWorld iterates over all of the objects in the world, and returns
// a sorted slice of the Intersections
func (w *World) IntersectWorld(r Ray) []*Intersection {
	xs := make([]*Intersection, 0, len(w.Objects))
	for _, s := range w.Objects {
		xs = append(xs, s.Intersect(r)...)
	}
	return Intersections(xs...)
}

// Given a set of pre computed values for an intersection, calculate the lighting/colour
// at that point in the world.
func (w *World) ShadeHit(c *Computations) Colour {
	return c.Object.GetMaterial().Lighting(
		c.Object,
		w.Light,
		c.Point,
		c.Eyev,
		c.Normalv,
		w.IsShadowed(c.OverPoint),
	)
}

func (w *World) ColourAt(r Ray) Colour {
	xs := w.IntersectWorld(r)
	hit := Hit(xs)
	if hit == nil {
		return NewColour(0, 0, 0)
	}
	c := hit.PrepareComputations(r)
	return w.ShadeHit(c)
}

func (w *World) IsShadowed(p Tuple) bool {
	v := w.Light.Position.Subtract(p)

	distance := v.Magnitude()
	direction := v.Normalise()
	r := NewRay(p, direction)
	xs := w.IntersectWorld(r)
	h := Hit(xs)
	if h != nil && h.T < distance {
		return true
	}
	return false
}
