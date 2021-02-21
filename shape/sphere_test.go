package shape_test

import (
	"testing"

	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/shape"
	"github.com/stretchr/testify/assert"
)

func TestASpheresHasADefaultTransformation(t *testing.T) {
	// Given
	s := shape.NewSphere()

	// Then
	assert.EqualValues(t, matrix.Identity(4), s.Transform)
}

func TestChangingASpheresTransformationMatrix(t *testing.T) {
	// Given
	s := shape.NewSphere()
	tm := matrix.Translation(2, 3, 4)

	// When
	s.SetTransform(tm)

	// Then
	assert.EqualValues(t, tm, s.Transform)
}
