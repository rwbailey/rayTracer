package tracer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rwbailey/ray/tracer"
)

func TestCreatingAndQueryingARay(t *testing.T) {
	// Given
	origin := Point(1, 2, 3)
	direction := Vector(4, 5, 6)

	// When
	r := NewRay(origin, direction)

	// Then
	assert.EqualValues(t, origin, r.Origin)
	assert.EqualValues(t, direction, r.Direction)
}

func TestComputingAPointFromADistance(t *testing.T) {
	// Given
	r := NewRay(Point(2, 3, 4), Vector(1, 0, 0))

	// Then
	assert.True(t, Point(2, 3, 4).Equals(r.Position(0)))
	assert.True(t, Point(3, 3, 4).Equals(r.Position(1)))
	assert.True(t, Point(1, 3, 4).Equals(r.Position(-1)))
	assert.True(t, Point(4.5, 3, 4).Equals(r.Position(2.5)))
}

func TestTranslatingARay(t *testing.T) {
	// Given
	r := NewRay(Point(1, 2, 3), Vector(0, 1, 0))
	m := TranslationMatrix(3, 4, 5)

	// When
	r2 := r.Transform(m)

	// Then
	assert.EqualValues(t, Point(4, 6, 8), r2.Origin)
	assert.EqualValues(t, Vector(0, 1, 0), r2.Direction)
}

func TestScalingARay(t *testing.T) {
	// Given
	r := NewRay(Point(1, 2, 3), Vector(0, 1, 0))
	m := ScalingMatrix(2, 3, 4)

	// When
	r2 := r.Transform(m)

	// Then
	assert.EqualValues(t, Point(2, 6, 12), r2.Origin)
	assert.EqualValues(t, Vector(0, 3, 0), r2.Direction)
}
