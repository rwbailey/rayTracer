package tuple_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rwbailey/ray/tuple"
)

func TestTupleCreation(t *testing.T) {
	// Given
	p := tuple.New(4.3, -4.2, 3.1, 1.0)
	v := tuple.New(4.3, -4.2, 3.1, 0.0)

	// Then
	assert.EqualValues(t, 4.3, p.X)
	assert.EqualValues(t, -4.2, p.Y)
	assert.EqualValues(t, 3.1, p.Z)
	assert.EqualValues(t, 1.0, p.W)
	assert.True(t, p.IsPoint())
	assert.False(t, p.IsVector())

	assert.EqualValues(t, 4.3, v.X)
	assert.EqualValues(t, -4.2, v.Y)
	assert.EqualValues(t, 3.1, v.Z)
	assert.EqualValues(t, 0.0, v.W)
	assert.True(t, v.IsVector())
	assert.False(t, v.IsPoint())
}

func TestTupleConstructors(t *testing.T) {
	// Given
	p := tuple.Point(4.3, -4.2, 3.1)
	v := tuple.Vector(4.3, -4.2, 3.1)

	// Then
	assert.EqualValues(t, tuple.Tuple{4.3, -4.2, 3.1, 1.0}, p)
	assert.EqualValues(t, tuple.Tuple{4.3, -4.2, 3.1, 0.0}, v)
}

func TestTupleComparison(t *testing.T) {
	// Given
	p1 := tuple.Point(1.3, 5.8, 8.6)
	p2 := tuple.Point(1.3, 5.8, 8.6)
	p3 := tuple.Point(56, 4.3, 12)

	v1 := tuple.Vector(1.3, 5.8, 8.6)
	v2 := tuple.Vector(1.3, 5.8, 8.6)
	v3 := tuple.Vector(56, 4.3, 12)

	// Then
	assert.True(t, p1.Equals(p2))
	assert.False(t, p1.Equals(p3))
	assert.True(t, v1.Equals(v2))
	assert.False(t, v1.Equals(v3))
	assert.False(t, p1.Equals(v1))
}

func TestTupleAddition(t *testing.T) {
	// Given
	p1 := tuple.Point(1, 2, 3)
	v1 := tuple.Vector(5, 5, 5)
	v2 := tuple.Vector(1, 2, 3)
	r1 := tuple.Point(6, 7, 8)
	r2 := tuple.Vector(6, 7, 8)

	// Then
	assert.EqualValues(t, r1, p1.Add(v1))
	assert.EqualValues(t, r2, v1.Add(v2))
}

func TestTupleSubtraction(t *testing.T) {
	// Given two points
	p1 := tuple.Point(1, 2, 3)
	p2 := tuple.Point(5, 5, 5)

	v1 := tuple.Vector(-4, -3, -2)

	// Then subtracting them yields the vector between them
	assert.EqualValues(t, v1, p1.Subtract(p2))

	// Given a point and a vector
	p3 := tuple.Point(1, 2, 3)
	v2 := tuple.Vector(5, 5, 5)

	p4 := tuple.Point(-4, -3, -2)

	// Then subtravting the vector from the point yields a new point
	assert.EqualValues(t, p4, p3.Subtract(v2))

	// Given two vectors
	v3 := tuple.Vector(1, 2, 3)
	v4 := tuple.Vector(5, 5, 5)

	v5 := tuple.Vector(-4, -3, -2)

	// Then subtracting them yields a new vector
	assert.EqualValues(t, v5, v3.Subtract(v4))
}

func TestTupleNegation(t *testing.T) {
	// Given
	v1 := tuple.Vector(5, 5, 5)
	v2 := tuple.Vector(-5, -5, -5)

	// Then
	assert.EqualValues(t, v2, v1.Negate())
}
