package light_test

import (
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestAPointLightHasAPositionAndIntensity(t *testing.T) {
	// Given
	intensity := colour.NewColour(1, 1, 1)
	position := tuple.Point(0, 0, 0)

	// When
	l := light.NewPointLight(position, intensity)

	// Then
	assert.EqualValues(t, intensity, l.Intensity)
	assert.EqualValues(t, position, l.Position)
}
