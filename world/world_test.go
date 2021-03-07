package world_test

import (
	"fmt"
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
	"github.com/rwbailey/ray/world"
	"github.com/stretchr/testify/assert"
)

func TestCreatingAnEmptyWorld(t *testing.T) {
	// Given
	w := world.New()

	// Then
	objects := w.Objects
	lgt := w.Light

	assert.EqualValues(t, 0, len(objects))
	assert.EqualValues(t, (*light.PointLight)(nil), lgt)
}

func TestCreateDefaultWorld(t *testing.T) {
	// Given
	ls := light.NewPointLight(tuple.Point(-10, 10, -10), colour.New(1, 1, 1))

	s1 := shape.NewSphere()
	s1.Material.Colour = colour.New(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := shape.NewSphere()
	s2.Transform = matrix.Scaling(0.5, 0.5, 0.5)

	// When
	w := world.Default()

	// Then
	assert.EqualValues(t, w.Light, ls)
	assert.Contains(t, w.Objects, s1)
	assert.Contains(t, w.Objects, s2)
}

func TestIntersectWorld(t *testing.T) {
	// Given
	w := world.Default()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))

	// When
	xs := w.IntersectWorld(r)

	// Then
	assert.EqualValues(t, 4, len(xs))
	assert.EqualValues(t, 4, xs[0].T)
	assert.EqualValues(t, 4.5, xs[1].T)
	assert.EqualValues(t, 5.5, xs[2].T)
	assert.EqualValues(t, 6, xs[3].T)
}

func TestShadeAnIntersection(t *testing.T) {
	// Given
	w := world.Default()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := w.Objects[0]
	i := &shape.Intersection{4, s}

	// When
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	fmt.Println(c)
	// Then
	assert.True(t, colour.New(0.38066, 0.47583, 0.2855).Equals(c))
}

func TestShadeAnIntersectionFromTheInside(t *testing.T) {
	// Given
	w := world.Default()
	w.Light = light.NewPointLight(tuple.Point(0, 0.25, 0), colour.New(1, 1, 1))
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := w.Objects[1]
	i := &shape.Intersection{0.5, s}

	// When
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	fmt.Println(c)

	// Then
	assert.True(t, colour.New(0.90498, 0.90498, 0.90498).Equals(c))
}
