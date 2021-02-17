package intersection_test

import (
	"testing"

	"github.com/rwbailey/ray/intersection"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestRaySphereIntersectionAtTwoPoints(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := intersection.Intersect(s, r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, 4.0, xs[0].T)
	assert.EqualValues(t, 6.0, xs[1].T)
}

func TestRaySphereIntersectionAtTangent(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 1, -5), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := intersection.Intersect(s, r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, 5.0, xs[0].T)
	assert.EqualValues(t, 5.0, xs[1].T)
}

func TestRaySphereNoIntersection(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 2, -5), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := intersection.Intersect(s, r)

	// Then
	assert.EqualValues(t, 0, len(xs))
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := intersection.Intersect(s, r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, -1.0, xs[0].T)
	assert.EqualValues(t, 1.0, xs[1].T)
}

func TestSphereBehindRay(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := intersection.Intersect(s, r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, -6.0, xs[0].T)
	assert.EqualValues(t, -4.0, xs[1].T)
}

func TestAnIntersectionEncapsulatesARayAndAnObject(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	i := intersection.Intersection{
		T:      3.5,
		Object: s,
	}

	// Then
	assert.EqualValues(t, 3.5, i.T)
	assert.EqualValues(t, s, i.Object)
}

func TestAggregatingIntersections(t *testing.T) {
	// Given
	s := shape.NewSphere()
	i1 := intersection.Intersection{1, s}
	i2 := intersection.Intersection{2, s}

	// When
	xs := intersection.Intersections(i1, i2)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, 1, xs[0].T)
	assert.EqualValues(t, 2, xs[1].T)
}

func TestIntersectSetsObject(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := intersection.Intersect(s, r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, s, xs[0].Object)
	assert.EqualValues(t, s, xs[1].Object)
}
