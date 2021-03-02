package world_test

import (
	"testing"

	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/world"
	"github.com/stretchr/testify/assert"
)

func TestCreatingAWorld(t *testing.T) {
	// Given
	w := world.New()

	// Then
	objects := w.Objects
	lgt := w.Light

	assert.EqualValues(t, 0, len(objects))
	assert.EqualValues(t, (*light.PointLight)(nil), lgt)
}
