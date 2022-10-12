package tracer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rwbailey/ray/tracer"
)

func TestTheNormalOfAPlaneIsTheSameEverywhere(t *testing.T) {
	// Given
	p := NewPlane()

	// When
	n1 := p.NormalAt(Point(0, 0, 0))
	n2 := p.NormalAt(Point(10, 0, -10))
	n3 := p.NormalAt(Point(-5, 0, 150))

	// Then
	N := Vector(0, 1, 0)
	assert.True(t, N.Equals(n1))
	assert.True(t, N.Equals(n2))
	assert.True(t, N.Equals(n3))
}

func TestIntersectWithARayParallelToThePlane(t *testing.T) {
	// Given
	p := NewPlane()
	r := NewRay(Point(0, 10, 0), Vector(0, 0, 1))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 0, len(xs))
}

func TestIntersectWithACoplanarRay(t *testing.T) {
	// Given
	p := NewPlane()
	r := NewRay(Point(0, 0, 0), Vector(0, 0, 1))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 0, len(xs))
}

func TestARayIntersectingAPlaneFromAbove(t *testing.T) {
	// Given
	p := NewPlane()
	r := NewRay(Point(0, 1, 0), Vector(0, -1, 0))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 1, len(xs))
	assert.Equal(t, float64(1), xs[0].T)
	assert.Equal(t, p, xs[0].Object)
}

func TestARayIntersectingAPlaneFromBelow(t *testing.T) {
	// Given
	p := NewPlane()
	r := NewRay(Point(0, -1, 0), Vector(0, 1, 0))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 1, len(xs))
	assert.Equal(t, float64(1), xs[0].T)
	assert.Equal(t, p, xs[0].Object)
}
