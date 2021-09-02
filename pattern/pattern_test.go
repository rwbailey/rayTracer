package pattern_test

import (
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/pattern"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestCreatingAStripePattern(t *testing.T) {
	// Given
	p := pattern.NewStripePattern(colour.White, colour.Black)

	// Then
	assert.Equal(t, colour.White, p.A)
	assert.Equal(t, colour.Black, p.B)
}

func TestAStripePatternIsConstantInY(t *testing.T) {
	// Given
	p := pattern.NewStripePattern(colour.White, colour.Black)

	// Then
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 0, 0)))
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 1, 0)))
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 2, 0)))
}

func TestAStripePatternIsConstantInZ(t *testing.T) {
	// Given
	p := pattern.NewStripePattern(colour.White, colour.Black)

	// Then
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 0, 0)))
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 0, 1)))
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 0, 2)))
}

func TestAStripePatternAlternatesInX(t *testing.T) {
	// Given
	p := pattern.NewStripePattern(colour.White, colour.Black)

	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0, 0, 0)))
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(0.9, 0, 0)))
	assert.Equal(t, colour.Black, p.ColourAt(tuple.Point(1, 0, 0)))
	assert.Equal(t, colour.Black, p.ColourAt(tuple.Point(-0.1, 0, 0)))
	assert.Equal(t, colour.Black, p.ColourAt(tuple.Point(-1, 0, 0)))
	assert.Equal(t, colour.White, p.ColourAt(tuple.Point(-1.1, 0, 0)))
}
