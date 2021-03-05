package shape_test

import (
	"testing"

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
	xs := s.Intersect(r)

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
	xs := s.Intersect(r)

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
	xs := s.Intersect(r)

	// Then
	assert.EqualValues(t, 0, len(xs))
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	s := shape.NewSphere()

	// When
	xs := s.Intersect(r)

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
	xs := s.Intersect(r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, -6.0, xs[0].T)
	assert.EqualValues(t, -4.0, xs[1].T)
}

func TestAnIntersectionEncapsulatesARayAndAnObject(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// When
	i := &shape.Intersection{
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
	i1 := &shape.Intersection{1, s}
	i2 := &shape.Intersection{2, s}

	// When
	xs := shape.Intersections(i1, i2)

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
	xs := s.Intersect(r)

	// Then
	assert.EqualValues(t, 2, len(xs))
	assert.EqualValues(t, s, xs[0].Object)
	assert.EqualValues(t, s, xs[1].Object)
}

func TestHitAllIntersectionsPositiveT(t *testing.T) {
	// Given
	s := shape.NewSphere()
	i1 := &shape.Intersection{1, s}
	i2 := &shape.Intersection{2, s}

	xs := shape.Intersections(i1, i2)

	// When
	i := shape.Hit(xs)

	// Then
	assert.EqualValues(t, i1, i)
}

func TestHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	// Given
	s := shape.NewSphere()
	i1 := &shape.Intersection{-1, s}
	i2 := &shape.Intersection{1, s}
	xs := shape.Intersections(i2, i1)

	// When
	i := shape.Hit(xs)

	// Then
	assert.EqualValues(t, i2, i)
}

func TestAllIntersectionsHaveNegativeT(t *testing.T) {
	// Given
	s := shape.NewSphere()
	i1 := &shape.Intersection{-2, s}
	i2 := &shape.Intersection{-1, s}
	xs := shape.Intersections(i2, i1)

	// When
	i := shape.Hit(xs)

	// Then
	var in *shape.Intersection
	assert.EqualValues(t, in, i)
	assert.True(t, in == nil)
}

func TestHitIsAlwaysLowestNonNegativeIntersection(t *testing.T) {
	// Given
	s := shape.NewSphere()
	i1 := &shape.Intersection{5, s}
	i2 := &shape.Intersection{7, s}
	i3 := &shape.Intersection{-3, s}
	i4 := &shape.Intersection{2, s}
	xs := shape.Intersections(i1, i2, i3, i4)

	// When
	i := shape.Hit(xs)

	// Then
	assert.EqualValues(t, i4, i)
}

// func TestPrecomputingTheStateOfAnIntersection(t *testing.T) {
// 	// Given
// 	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
// 	s := shape.NewSphere()
// 	i := shape.Intersection{4, s}

// 	// When
// 	comps := i.PrepareComputations(r)

// 	// Then
// 	assert.EqualValues(t, i.T, comps.T)
// 	assert.EqualValues(t, i.Object, comps.Object)
// 	assert.EqualValues(t, tuple.Point(0, 0, -1), comps.Point)
// 	assert.EqualValues(t, tuple.Vector(0, 0, -1), comps.Eyev)
// 	assert.EqualValues(t, tuple.Vector(0, 0, -1), comps.Normalv)
// }
