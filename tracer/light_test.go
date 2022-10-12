package tracer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rwbailey/ray/tracer"
)

func TestAPointLightHasAPositionAndIntensity(t *testing.T) {
	// Given
	intensity := NewColour(1, 1, 1)
	position := Point(0, 0, 0)

	// When
	l := NewPointLight(position, intensity)

	// Then
	assert.EqualValues(t, intensity, l.Intensity)
	assert.EqualValues(t, position, l.Position)
}
