package shape_test

import (
	"testing"

	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestASpheresHasADefaultTransformation(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// Then
	assert.EqualValues(t, matrix.Identity(4), s.Transform)
}

func TestChangingASpheresTransformationMatrix(t *testing.T) {
	// Given
	s := shape.NewSphere()
	tm := matrix.Translation(2, 3, 4)

	// When
	s.SetTransform(tm)

	// Then
	assert.EqualValues(t, tm, s.Transform)
}

func TestIntersectingAScaledSphereWithAnArray(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()
	tm := matrix.Scaling(2, 2, 2)

	// When
	s.SetTransform(tm)
	xs := s.Intersect(r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, 3, xs[0].T)
	assert.EqualValues(t, 7, xs[1].T)
}
