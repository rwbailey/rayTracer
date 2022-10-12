package camera_test

import (
	"math"
	"testing"

	"github.com/rwbailey/ray/camera"
	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/tuple"
	"github.com/rwbailey/ray/world"
	"github.com/stretchr/testify/assert"
)

func TestConstructingACamera(t *testing.T) {
	// Given
	hsize := 160.0
	vsize := 120.0
	fieldOfView := math.Pi / 2

	// When
	c := camera.NewCamera(hsize, vsize, fieldOfView)

	// Then
	assert.EqualValues(t, 160, c.HSize)
	assert.EqualValues(t, 120, c.VSize)
	assert.EqualValues(t, math.Pi/2, c.FieldOfView)
	assert.EqualValues(t, matrix.IdentityMatrix(4), c.Transform)
}

func TestThePixelSizeForAHorizontalCanvas(t *testing.T) {
	// Given
	c := camera.NewCamera(200, 125, math.Pi/2)

	// Then
	assert.EqualValues(t, 0.01, c.PixelSize)
}

func TestThePixelSizeForAVerticalCanvas(t *testing.T) {
	// Given
	c := camera.NewCamera(125, 200, math.Pi/2)

	// Then
	assert.EqualValues(t, 0.01, c.PixelSize)
}

func TestConstructingARayThroughTheCentreOfTheCanvas(t *testing.T) {
	// Given
	c := camera.NewCamera(201, 101, math.Pi/2)

	// When
	r := c.RayForPixel(100, 50)

	// Then
	assert.Equal(t, tuple.Point(0, 0, 0), r.Origin)
	assert.True(t, tuple.Vector(0, 0, -1).Equals(r.Direction))
}

func TestConstructingARayThroughTheCornerOfTheCanvas(t *testing.T) {
	// Given
	c := camera.NewCamera(201, 101, math.Pi/2)

	// When
	r := c.RayForPixel(0, 0)

	// Then
	assert.Equal(t, tuple.Point(0, 0, 0), r.Origin)
	assert.True(t, tuple.Vector(0.66519, 0.33259, -0.66851).Equals(r.Direction))
}

func TestConstructingARayWhenTheCameraIsTransformed(t *testing.T) {
	// Given
	c := camera.NewCamera(201, 101, math.Pi/2)

	// When
	c.Transform = matrix.RotationYMatrix(math.Pi / 4).MultiplyMatrix(matrix.TranslationMatrix(0, -2, 5))
	r := c.RayForPixel(100, 50)

	// Then
	assert.Equal(t, tuple.Point(0, 2, -5), r.Origin)
	assert.True(t, tuple.Vector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2).Equals(r.Direction))
}

func TestRenderingAWorldWithACamera(t *testing.T) {
	// Given
	w := world.DefaultWorld()
	c := camera.NewCamera(11, 11, math.Pi/2)
	from := tuple.Point(0, 0, -5)
	to := tuple.Point(0, 0, 0)
	up := tuple.Vector(0, 1, 0)
	c.Transform = matrix.ViewTransform(from, to, up)

	// When
	img := c.Render(w)

	// Then
	assert.True(t, colour.NewColour(0.38066, 0.47583, 0.2855).Equals(img.PixelAt(5, 5)))
}
