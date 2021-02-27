package material_test

import (
	"testing"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/material"
	"github.com/stretchr/testify/assert"
)

func TestDefaultMaterial(t *testing.T) {
	// Given
	m := material.New()

	// Then
	assert.EqualValues(t, colour.New(1, 1, 1), m.Colour)
	assert.EqualValues(t, 0.1, m.Ambient)
	assert.EqualValues(t, 0.9, m.Diffuse)
	assert.EqualValues(t, 0.9, m.Specular)
	assert.EqualValues(t, 200.0, m.Shininess)
}
