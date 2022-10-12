package material_test

import (
	"math"
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/material"
	"github.com/rwbailey/ray/pattern"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

var background = material.NewMaterial()
var point = tuple.Point(0, 0, 0)

func TestDefaultMaterial(t *testing.T) {
	// Given
	m := material.NewMaterial()

	// Then
	assert.EqualValues(t, colour.NewColour(1, 1, 1), m.Colour)
	assert.EqualValues(t, 0.1, m.Ambient)
	assert.EqualValues(t, 0.9, m.Diffuse)
	assert.EqualValues(t, 0.9, m.Specular)
	assert.EqualValues(t, 200.0, m.Shininess)
}

func TestLightingWithTheEyeBetweenTheLightAndTheSurface(t *testing.T) {
	// Given
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 0, -10), colour.NewColour(1, 1, 1))

	// When
	result := background.Lighting(light, point, eyev, normalv, false)

	// Then
	assert.EqualValues(t, colour.NewColour(1.9, 1.9, 1.9), result)
}

func TestLightingWithTheEyeBetweenTheLightAndTheSurface45Offset(t *testing.T) {
	// Given
	eyev := tuple.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 0, -10), colour.NewColour(1, 1, 1))

	// When
	result := background.Lighting(light, point, eyev, normalv, false)

	// Then
	assert.EqualValues(t, colour.NewColour(1.0, 1.0, 1.0), result)
}

func TestLightingWithTheOppositeSurfaceAndLightAt45Offset(t *testing.T) {
	// Given
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 10, -10), colour.NewColour(1, 1, 1))

	// When
	result := background.Lighting(light, point, eyev, normalv, false)

	// Then
	assert.True(t, colour.NewColour(0.7364, 0.7364, 0.7364).Equals(result))
}

func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	// Given
	eyev := tuple.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 10, -10), colour.NewColour(1, 1, 1))

	// When
	result := background.Lighting(light, point, eyev, normalv, false)

	// Then
	assert.True(t, colour.NewColour(1.6364, 1.6364, 1.6364).Equals(result))
}

func TestLightingWithLightBehindTheSurface(t *testing.T) {
	// Given
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 0, 10), colour.NewColour(1, 1, 1))

	// When
	result := background.Lighting(light, point, eyev, normalv, false)

	// Then
	assert.EqualValues(t, colour.NewColour(0.1, 0.1, 0.1), result)
}

func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	// Given
	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 0, -10), colour.NewColour(1, 1, 1))
	inShadow := true

	// When
	result := background.Lighting(light, point, eyev, normalv, inShadow)

	// Then
	assert.Equal(t, colour.NewColour(0.1, 0.1, 0.1), result)
}

func TestLightingWithAPatternApplied(t *testing.T) {
	// Given
	background.Pattern = pattern.NewStripePattern(colour.White, colour.Black)
	background.Ambient = 1
	background.Diffuse = 0
	background.Specular = 0

	eyev := tuple.Vector(0, 0, -1)
	normalv := tuple.Vector(0, 0, -1)
	light := light.NewPointLight(tuple.Point(0, 0, -10), colour.White)

	// When
	c1 := background.Lighting(light, tuple.Point(0.9, 0, 0), eyev, normalv, false)
	c2 := background.Lighting(light, tuple.Point(1.1, 0, 0), eyev, normalv, false)

	// Then
	assert.Equal(t, colour.White, c1)
	assert.Equal(t, colour.Black, c2)
}
