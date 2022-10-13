package tracer_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rwbailey/ray/tracer"
)

var obj = NewSphere()
var point = Point(0, 0, 0)

func TestDefaultMaterial(t *testing.T) {
	// Given
	m := NewMaterial()

	// Then
	assert.EqualValues(t, NewColour(1, 1, 1), m.Colour)
	assert.EqualValues(t, 0.1, m.Ambient)
	assert.EqualValues(t, 0.9, m.Diffuse)
	assert.EqualValues(t, 0.9, m.Specular)
	assert.EqualValues(t, 200.0, m.Shininess)
}

func TestLightingWithTheEyeBetweenTheLightAndTheSurface(t *testing.T) {
	// Given
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), NewColour(1, 1, 1))

	obj.SetMaterial(NewMaterial())

	// When
	result := obj.GetMaterial().Lighting(obj, light, point, eyev, normalv, false)

	// Then
	assert.EqualValues(t, NewColour(1.9, 1.9, 1.9), result)
}

func TestLightingWithTheEyeBetweenTheLightAndTheSurface45Offset(t *testing.T) {
	// Given
	eyev := Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), NewColour(1, 1, 1))

	obj.SetMaterial(NewMaterial())

	// When
	result := obj.GetMaterial().Lighting(obj, light, point, eyev, normalv, false)

	// Then
	assert.EqualValues(t, NewColour(1.0, 1.0, 1.0), result)
}

func TestLightingWithTheOppositeSurfaceAndLightAt45Offset(t *testing.T) {
	// Given
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 10, -10), NewColour(1, 1, 1))

	obj.SetMaterial(NewMaterial())

	// When
	result := obj.GetMaterial().Lighting(obj, light, point, eyev, normalv, false)

	// Then
	assert.True(t, NewColour(0.7364, 0.7364, 0.7364).Equals(result))
}

func TestLightingWithEyeInPathOfReflectionVector(t *testing.T) {
	// Given
	eyev := Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 10, -10), NewColour(1, 1, 1))

	obj.SetMaterial(NewMaterial())

	// When
	result := obj.GetMaterial().Lighting(obj, light, point, eyev, normalv, false)

	// Then
	assert.True(t, NewColour(1.6364, 1.6364, 1.6364).Equals(result))
}

func TestLightingWithLightBehindTheSurface(t *testing.T) {
	// Given
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, 10), NewColour(1, 1, 1))

	obj.SetMaterial(NewMaterial())

	// When
	result := obj.GetMaterial().Lighting(obj, light, point, eyev, normalv, false)

	// Then
	assert.EqualValues(t, NewColour(0.1, 0.1, 0.1), result)
}

func TestLightingWithTheSurfaceInShadow(t *testing.T) {
	// Given
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), NewColour(1, 1, 1))
	inShadow := true

	obj.SetMaterial(NewMaterial())

	// When
	result := obj.GetMaterial().Lighting(obj, light, point, eyev, normalv, inShadow)

	// Then
	assert.Equal(t, NewColour(0.1, 0.1, 0.1), result)
}

func TestLightingWithAPatternApplied(t *testing.T) {
	// Given
	obj.SetMaterial(NewMaterial())
	obj.GetMaterial().Pattern = NewStripePattern(White, Black)
	obj.GetMaterial().Ambient = 1
	obj.GetMaterial().Diffuse = 0
	obj.GetMaterial().Specular = 0

	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), White)

	// When
	c1 := obj.GetMaterial().Lighting(obj, light, Point(0.9, 0, 0), eyev, normalv, false)
	c2 := obj.GetMaterial().Lighting(obj, light, Point(1.1, 0, 0), eyev, normalv, false)

	// Then
	assert.Equal(t, White, c1)
	assert.Equal(t, Black, c2)
}
