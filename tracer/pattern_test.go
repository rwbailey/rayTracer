package tracer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/rwbailey/ray/tracer"
)

func TestCreatingAStripePattern(t *testing.T) {
	// Given
	p := NewStripePattern(White, Black)

	// Then
	assert.Equal(t, White, p.A)
	assert.Equal(t, Black, p.B)
}

func TestAStripePatternIsConstantInY(t *testing.T) {
	// Given
	p := NewStripePattern(White, Black)

	// Then
	assert.Equal(t, White, p.ColourAt(Point(0, 0, 0)))
	assert.Equal(t, White, p.ColourAt(Point(0, 1, 0)))
	assert.Equal(t, White, p.ColourAt(Point(0, 2, 0)))
}

func TestAStripePatternIsConstantInZ(t *testing.T) {
	// Given
	p := NewStripePattern(White, Black)

	// Then
	assert.Equal(t, White, p.ColourAt(Point(0, 0, 0)))
	assert.Equal(t, White, p.ColourAt(Point(0, 0, 1)))
	assert.Equal(t, White, p.ColourAt(Point(0, 0, 2)))
}

func TestAStripePatternAlternatesInX(t *testing.T) {
	// Given
	p := NewStripePattern(White, Black)

	assert.Equal(t, White, p.ColourAt(Point(0, 0, 0)))
	assert.Equal(t, White, p.ColourAt(Point(0.9, 0, 0)))
	assert.Equal(t, Black, p.ColourAt(Point(1, 0, 0)))
	assert.Equal(t, Black, p.ColourAt(Point(-0.1, 0, 0)))
	assert.Equal(t, Black, p.ColourAt(Point(-1, 0, 0)))
	assert.Equal(t, White, p.ColourAt(Point(-1.1, 0, 0)))
}

func TestLightingWithPattern(t *testing.T) {
	// Given
	m := NewMaterial()
	m.Pattern = NewStripePattern(White, Black)
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyev := Vector(0, 0, -1)
	normalv := Vector(0, 0, -1)
	light := NewPointLight(Point(0, 0, -10), White)

	// When
	c1 := m.Lighting(light, Point(0.9, 0, 0), eyev, normalv, false)
	c2 := m.Lighting(light, Point(1.1, 0, 0), eyev, normalv, false)

	// Then
	assert.Equal(t, White, c1)
	assert.Equal(t, Black, c2)
}

func TestStripesWithObjectTransformation(t *testing.T) {
	// Given
	obj := NewSphere()
	obj.SetTransform(ScalingMatrix(2, 2, 2))
	obj.Material.Pattern = NewStripePattern(White, Black)

	// When
	c := obj.Material.Pattern.ColourAtObject(obj, Point(1.5, 0, 0))

	// Then
	assert.Equal(t, White, c)
}

func TestStripesWithPatternTransformation(t *testing.T) {
	// Given
	obj := NewSphere()
	obj.Material.Pattern = NewStripePattern(White, Black)
	obj.Material.Pattern.SetTransform(ScalingMatrix(2, 2, 2))

	// When
	c := obj.Material.Pattern.ColourAtObject(obj, Point(1.5, 0, 0))

	// Then
	assert.Equal(t, c, White)
}

func TestStripesWithPatternAndObjectTransformations(t *testing.T) {
	// Given
	obj := NewSphere()
	obj.SetTransform(ScalingMatrix(2, 2, 2))
	obj.Material.Pattern = NewStripePattern(White, Black)
	obj.Material.Pattern.SetTransform(TranslationMatrix(0.5, 0, 0))

	// When
	c := obj.Material.Pattern.ColourAtObject(obj, Point(2.5, 0, 0))

	// Then
	assert.Equal(t, White, c)
}
