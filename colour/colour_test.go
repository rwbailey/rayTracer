package colour_test

import (
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/stretchr/testify/assert"
)

func TestCreateColour(t *testing.T) {
	// Given
	c := colour.New(-0.5, 0.4, 1.7)

	// Then
	assert.EqualValues(t, -0.5, c.Red)
	assert.EqualValues(t, 0.4, c.Green)
	assert.EqualValues(t, 1.7, c.Blue)
}
