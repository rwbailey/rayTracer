package world_test

import (
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/matrix"
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
