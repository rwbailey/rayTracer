package ray_test

import (
	"testing"

	"github.com/rwbailey/ray/matrix"
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

func TestTranslatingARay(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	m := matrix.Translation(3, 4, 5)

	// When
	r2 := r.Transform(m)

	// Then
	assert.EqualValues(t, tuple.Point(4, 6, 8), r2.Origin)
	assert.EqualValues(t, tuple.Vector(0, 1, 0), r2.Direction)
}

func TestScalingARay(t *testing.T) {
	// Given
	r := ray.New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	m := matrix.Scaling(2, 3, 4)

	// When
	r2 := r.Transform(m)

	// Then
	assert.EqualValues(t, tuple.Point(2, 6, 12), r2.Origin)
	assert.EqualValues(t, tuple.Vector(0, 3, 0), r2.Direction)
}
