package pattern_test

import (
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/material"
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

func TestLightingWithPattern(t *testing.T) {
	// Given
	m := material.New()
	m.Pattern = pattern.NewStripePattern(colour.White, colour.Black)
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 0, -10), colour.White)

	// When
	c1 := m.Lighting(light, tuple.Point(0.9, 0, 0), eyev, normalv, false)
	c2 := m.Lighting(light, tuple.Point(1.1, 0, 0), eyev, normalv, false)

	// Then
	assert.Equal(t, colour.White, c1)
	assert.Equal(t, colour.Black, c2)
}
