package camera_test

import (
	"math"
	"testing"

	"github.com/rwbailey/ray/camera"
	"github.com/rwbailey/ray/matrix"
	"github.com/stretchr/testify/assert"
)

func TestConstructingACamera(t *testing.T) {
	// Given
	hsize := 160.0
	vsize := 120.0
	fieldOfView := math.Pi / 2

	// When
	c := camera.New(hsize, vsize, fieldOfView)

	// Then
	assert.EqualValues(t, 160, c.HSize)
	assert.EqualValues(t, 120, c.VSize)
	assert.EqualValues(t, math.Pi/2, c.FieldOfView)
	assert.EqualValues(t, matrix.Identity(4), c.Transform)
}

func TestThePixelSizeForAHorizontalCanvas(t *testing.T) {
	// Given
	c := camera.New(200, 125, math.Pi/2)

	// Then
	assert.EqualValues(t, 0.01, c.PixelSize)
}

func TestThePixelSizeForAVerticalCanvas(t *testing.T) {
	// Given
	c := camera.New(125, 200, math.Pi/2)

	// Then
	assert.EqualValues(t, 0.01, c.PixelSize)
}
