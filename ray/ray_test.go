package ray_test

import (
	"testing"

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
