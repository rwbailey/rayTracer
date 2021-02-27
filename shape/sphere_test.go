package shape_test

import (
	"math"
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

func TestTheNormalAtAPointOnTheXAxis(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	n := s.NormalAt(tuple.Point(1, 0, 0))

	// Then
	assert.True(t, tuple.Vector(1, 0, 0).Equals(n))
}

func TestTheNormalAtAPointOnTheYAxis(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	n := s.NormalAt(tuple.Point(0, 1, 0))

	// Then
	assert.True(t, tuple.Vector(0, 1, 0).Equals(n))
}

func TestTheNormalAtAPointOnTheZAxis(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	n := s.NormalAt(tuple.Point(0, 0, 1))

	// Then
	assert.True(t, tuple.Vector(0, 0, 1).Equals(n))
}

func TestTheNormalAtANonAxialPoint(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	n := s.NormalAt(tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	// Then
	assert.True(t, tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Equals(n))
}

func TestTheNormalIsNormalised(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	n := s.NormalAt(tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	// Then
	assert.True(t, n.Normalise().Equals(n))
}

func TestNormalOnATranslatedSphere(t *testing.T) {
	// Given
	s := shape.NewSphere()
	s.SetTransform(matrix.Translation(0, 1, 0))

	// When
	n := s.NormalAt(tuple.Point(0, 1.70711, -0.70711))

	// Then
	assert.True(t, n.Equals(tuple.Vector(0, 0.70711, -0.70711)))
}

func TestNormalOnATransformedSphere(t *testing.T) {
	// Given
	s := shape.NewSphere()
	m := matrix.Identity(4).RotateZ(math.Pi/5).Scale(1, 0.5, 1)
	s.SetTransform(m)

	// When
	n := s.NormalAt(tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	// Then
	assert.True(t, n.Equals(tuple.Vector(0, 0.97014, -0.24254)))
}
