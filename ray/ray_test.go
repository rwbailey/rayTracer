package ray_test

import (
	"testing"

	"github.com/rwbailey/ray/object"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestCreatingAndQueryingARay(t *testing.T) {
	// Given
	origin := tuple.Point(1, 2, 3)
	direction := tuple.Vector(4, 5, 6)

	// When
	r := ray.New(origin, direction)

	// Then
	assert.EqualValues(t, origin, r.Origin)
	assert.EqualValues(t, direction, r.Direction)
}

func TestComputingAPointFromADistance(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(2, 3, 4), tuple.Vector(1, 0, 0))

	// Then
	assert.True(t, tuple.Point(2, 3, 4).Equals(r.Position(0)))
	assert.True(t, tuple.Point(3, 3, 4).Equals(r.Position(1)))
	assert.True(t, tuple.Point(1, 3, 4).Equals(r.Position(-1)))
	assert.True(t, tuple.Point(4.5, 3, 4).Equals(r.Position(2.5)))
}

func TestRaySphereIntersectionAtTwoPoints(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := object.NewSphere()

	// When
	xs := r.Intersect(s)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, 4.0, xs[0])
	assert.EqualValues(t, 6.0, xs[1])
}

func TestRaySphereIntersectionAtTangent(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 1, -5), tuple.Vector(0, 0, 1))
	s := object.NewSphere()

	// When
	xs := r.Intersect(s)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, 5.0, xs[0])
	assert.EqualValues(t, 5.0, xs[1])
}

func TestRaySphereNoIntersection(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 2, -5), tuple.Vector(0, 0, 1))
	s := object.NewSphere()

	// When
	xs := r.Intersect(s)

	// Then
	assert.EqualValues(t, 0, len(xs))
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := object.NewSphere()

	// When
	xs := r.Intersect(s)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, -1.0, xs[0])
	assert.EqualValues(t, 1.0, xs[1])
}

func TestSphereBehindRay(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
	s := object.NewSphere()

	// When
	xs := r.Intersect(s)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, -6.0, xs[0])
	assert.EqualValues(t, -4.0, xs[1])
}
