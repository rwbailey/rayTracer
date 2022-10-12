package tracer_test

import (
	"fmt"
	"testing"

	"github.com/rwbailey/ray/tracer"

	"github.com/stretchr/testify/assert"
)

func TestCreatingAnEmptyWorld(t *testing.T) {
	// Given
	w := tracer.NewWorld()

	// Then
	objects := w.Objects
	lgt := w.Light

	assert.EqualValues(t, 0, len(objects))
	assert.EqualValues(t, (*tracer.PointLight)(nil), lgt)
}

func TestCreateDefaultWorld(t *testing.T) {
	// Given
	ls := tracer.NewPointLight(tracer.Point(-10, 10, -10), tracer.NewColour(1, 1, 1))

	s1 := tracer.NewSphere()
	s1.Material.Colour = tracer.NewColour(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := tracer.NewSphere()
	s2.Transform = tracer.ScalingMatrix(0.5, 0.5, 0.5)

	// When
	w := tracer.DefaultWorld()

	// Then
	assert.EqualValues(t, w.Light, ls)
	assert.Contains(t, w.Objects, s1)
	assert.Contains(t, w.Objects, s2)
}

func TestIntersectWorld(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	r := tracer.NewRay(tracer.Point(0, 0, -5), tracer.Vector(0, 0, 1))

	// When
	xs := w.IntersectWorld(r)

	// Then
	assert.EqualValues(t, 4, len(xs))
	assert.EqualValues(t, 4, xs[0].T)
	assert.EqualValues(t, 4.5, xs[1].T)
	assert.EqualValues(t, 5.5, xs[2].T)
	assert.EqualValues(t, 6, xs[3].T)
}

func TestShadeAnIntersection(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	r := tracer.NewRay(tracer.Point(0, 0, -5), tracer.Vector(0, 0, 1))
	s := w.Objects[0]
	i := &tracer.Intersection{T: 4, Object: s}

	// When
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	fmt.Println(c)
	// Then
	assert.True(t, tracer.NewColour(0.38066, 0.47583, 0.2855).Equals(c))
}

func TestShadeAnIntersectionFromTheInside(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	w.Light = tracer.NewPointLight(tracer.Point(0, 0.25, 0), tracer.NewColour(1, 1, 1))
	r := tracer.NewRay(tracer.Point(0, 0, 0), tracer.Vector(0, 0, 1))
	s := w.Objects[1]
	i := &tracer.Intersection{T: 0.5, Object: s}

	// When
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	fmt.Println(c)

	// Then
	assert.True(t, tracer.NewColour(0.90498, 0.90498, 0.90498).Equals(c))
}

func TestTheColourWhenARayMisses(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	r := tracer.NewRay(tracer.Point(0, 0, -5), tracer.Vector(0, 1, 0))

	// When
	c := w.ColourAt(r)

	// Then
	assert.True(t, tracer.NewColour(0, 0, 0).Equals(c))
}

func TestTheColourWhenARayHits(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	r := tracer.NewRay(tracer.Point(0, 0, -5), tracer.Vector(0, 0, 1))

	// When
	c := w.ColourAt(r)

	// Then
	assert.True(t, tracer.NewColour(0.38066, 0.47583, 0.2855).Equals(c))

}

func TestTheColourWithAnIntersectionBehindTheRay(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	outer := w.Objects[0]
	outer.GetMaterial().Ambient = 1
	inner := w.Objects[1]
	inner.GetMaterial().Ambient = 1
	r := tracer.NewRay(tracer.Point(0, 0, 0.75), tracer.Vector(0, 0, -1))

	// When
	c := w.ColourAt(r)

	// Then
	assert.True(t, inner.GetMaterial().Colour.Equals(c))
}

func TestThereIsNoShadowWhenNothingIsColinearWithThePointAndLight(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	p := tracer.Point(0, 10, 0)

	// Then
	assert.False(t, w.IsShadowed(p))
}

func TestTheShadowWhenAnObjectIsBetweenThePointAndTheLIght(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	p := tracer.Point(10, -10, 10)

	// Then
	assert.True(t, w.IsShadowed(p))
}

func TestTheShadowWhenAnObjectIsBehindTheLight(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	p := tracer.Point(-20, 20, -20)

	// Then
	assert.False(t, w.IsShadowed(p))
}

func TestTheShadowWhenAnObjectIsBehindThePoint(t *testing.T) {
	// Given
	w := tracer.DefaultWorld()
	p := tracer.Point(-2, 2, -2)

	// Then
	assert.False(t, w.IsShadowed(p))
}

func TestShadeHitIsGivenAnIntersectionInShadow(t *testing.T) {
	// Given
	w := tracer.NewWorld()
	w.Light = tracer.NewPointLight(tracer.Point(0, 0, -10), tracer.White)
	s1 := tracer.NewSphere()
	w.Objects = append(w.Objects, s1)
	s2 := tracer.NewSphere()
	s2.Transform = tracer.TranslationMatrix(0, 0, 10)
	w.Objects = append(w.Objects, s2)
	r := tracer.NewRay(tracer.Point(0, 0, 5), tracer.Vector(0, 0, 1))
	i := &tracer.Intersection{T: 4, Object: s2}

	// When
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)

	// Then
	assert.True(t, tracer.NewColour(0.1, 0.1, 0.1).Equals(c))
}

func TestTheHitShouldOffsetThePoint(t *testing.T) {
	// Given
	r := tracer.NewRay(tracer.Point(0, 0, -5), tracer.Vector(0, 0, 1))
	s := tracer.NewSphere()
	s.Transform = tracer.TranslationMatrix(0, 0, 1)
	i := &tracer.Intersection{T: 5, Object: s}

	// When
	comps := i.PrepareComputations(r)

	// Then
	assert.Less(t, comps.OverPoint.Z, -tracer.Epsilon/2)
	assert.Greater(t, comps.Point.Z, comps.OverPoint.Z)
}
