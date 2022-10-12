package shape_test

import (
	"testing"

	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestTheNormalOfAPlaneIsTheSameEverywhere(t *testing.T) {
	// Given
	p := shape.NewPlane()

	// When
	n1 := p.NormalAt(tuple.Point(0, 0, 0))
	n2 := p.NormalAt(tuple.Point(10, 0, -10))
	n3 := p.NormalAt(tuple.Point(-5, 0, 150))

	// Then
	N := tuple.Vector(0, 1, 0)
	assert.True(t, N.Equals(n1))
	assert.True(t, N.Equals(n2))
	assert.True(t, N.Equals(n3))
}

func TestIntersectWithARayParallelToThePlane(t *testing.T) {
	// Given
	p := shape.NewPlane()
	r := ray.NewRay(tuple.Point(0, 10, 0), tuple.Vector(0, 0, 1))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 0, len(xs))
}

func TestIntersectWithACoplanarRay(t *testing.T) {
	// Given
	p := shape.NewPlane()
	r := ray.NewRay(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 0, len(xs))
}

func TestARayIntersectingAPlaneFromAbove(t *testing.T) {
	// Given
	p := shape.NewPlane()
	r := ray.NewRay(tuple.Point(0, 1, 0), tuple.Vector(0, -1, 0))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 1, len(xs))
	assert.Equal(t, float64(1), xs[0].T)
	assert.Equal(t, p, xs[0].Object)
}

func TestARayIntersectingAPlaneFromBelow(t *testing.T) {
	// Given
	p := shape.NewPlane()
	r := ray.NewRay(tuple.Point(0, -1, 0), tuple.Vector(0, 1, 0))

	// When
	xs := p.Intersect(r)

	// Then
	assert.Equal(t, 1, len(xs))
	assert.Equal(t, float64(1), xs[0].T)
	assert.Equal(t, p, xs[0].Object)
}
