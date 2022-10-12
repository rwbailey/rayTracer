package tracer_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rwbailey/ray/tracer"
)

func TestASpheresHasADefaultTransformation(t *testing.T) {
	// Given
	s := NewSphere()

	// Then
	assert.EqualValues(t, IdentityMatrix(4), s.Transform)
}

func TestChangingASpheresTransformationMatrix(t *testing.T) {
	// Given
	s := NewSphere()
	tm := TranslationMatrix(2, 3, 4)

	// When
	s.SetTransform(tm)

	// Then
	assert.EqualValues(t, tm, s.Transform)
}

func TestIntersectingAScaledSphereWithAnArray(t *testing.T) {
	// Given
	r := NewRay(Point(0, 0, -5), Vector(0, 0, 1))
	s := NewSphere()
	tm := ScalingMatrix(2, 2, 2)

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
	s := NewSphere()

	// When
	n := s.NormalAt(Point(1, 0, 0))

	// Then
	assert.True(t, Vector(1, 0, 0).Equals(n))
}

func TestTheNormalAtAPointOnTheYAxis(t *testing.T) {
	// Given
	s := NewSphere()

	// When
	n := s.NormalAt(Point(0, 1, 0))

	// Then
	assert.True(t, Vector(0, 1, 0).Equals(n))
}

func TestTheNormalAtAPointOnTheZAxis(t *testing.T) {
	// Given
	s := NewSphere()

	// When
	n := s.NormalAt(Point(0, 0, 1))

	// Then
	assert.True(t, Vector(0, 0, 1).Equals(n))
}

func TestTheNormalAtANonAxialPoint(t *testing.T) {
	// Given
	s := NewSphere()

	// When
	n := s.NormalAt(Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	// Then
	assert.True(t, Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Equals(n))
}

func TestTheNormalIsNormalised(t *testing.T) {
	// Given
	s := NewSphere()

	// When
	n := s.NormalAt(Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	// Then
	assert.True(t, n.Normalise().Equals(n))
}

func TestNormalOnATranslatedSphere(t *testing.T) {
	// Given
	s := NewSphere()
	s.SetTransform(TranslationMatrix(0, 1, 0))

	// When
	n := s.NormalAt(Point(0, 1.70711, -0.70711))

	// Then
	assert.True(t, n.Equals(Vector(0, 0.70711, -0.70711)))
}

func TestNormalOnATransformedSphere(t *testing.T) {
	// Given
	s := NewSphere()
	m := IdentityMatrix(4).RotateZ(math.Pi/5).Scale(1, 0.5, 1)
	s.SetTransform(m)

	// When
	n := s.NormalAt(Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	// Then
	assert.True(t, n.Equals(Vector(0, 0.97014, -0.24254)))
}

func TestASphereHasADefaultMaterial(t *testing.T) {
	// Given
	s := NewSphere()

	// Then
	assert.EqualValues(t, NewMaterial(), s.Material)
}

func TestASphereMayBeAssignedAMaterial(t *testing.T) {
	// Given
	s := NewSphere()
	m := NewMaterial()
	m.Ambient = 1

	// When
	s.Material = m

	// Then
	assert.EqualValues(t, m, s.Material)
}
