package colour_test

import (
	"testing"

	. "github.com/rwbailey/ray/colour"
	"github.com/stretchr/testify/assert"
)

func TestColourCreation(t *testing.T) {
	// Given
	c := NewColour(-0.5, 0.4, 1.7)

	// Then
	assert.EqualValues(t, -0.5, c.Red)
	assert.EqualValues(t, 0.4, c.Green)
	assert.EqualValues(t, 1.7, c.Blue)
}

func TestColourAddition(t *testing.T) {
	// Given
	c1 := NewColour(0.9, 0.6, 0.75)
	c2 := NewColour(0.7, 0.1, 0.25)

	// Then
	assert.EqualValues(t, NewColour(1.6, 0.7, 1.0), c1.Add(c2))
}

func TestColourSubtraction(t *testing.T) {
	// Given
	c1 := NewColour(0.9, 0.6, 0.75)
	c2 := NewColour(0.7, 0.1, 0.25)

	c3 := NewColour(0.2, 0.5, 0.5)

	// Then
	assert.True(t, c3.Equals(c1.Subtract(c2)))
}

func TestColourMultiplicationByScalar(t *testing.T) {
	// Given
	c := NewColour(0.2, 0.3, 0.4)

	// Then
	assert.EqualValues(t, NewColour(0.4, 0.6, 0.8), c.Multiply(2.0))
}

func TestColourMultiplicationByColour(t *testing.T) {
	// Given
	a := NewColour(1.0, 0.2, 0.4)
	b := NewColour(0.9, 1, 0.1)

	// Then
	assert.True(t, NewColour(0.9, 0.2, 0.04).Equals(a.Product(b)))
}
