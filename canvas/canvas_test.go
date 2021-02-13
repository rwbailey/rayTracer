package canvas_test

import (
	"strings"
	"testing"

	"github.com/rwbailey/ray/canvas"
	"github.com/rwbailey/ray/colour"
	"github.com/stretchr/testify/assert"
)

func TestCanvasCreation(t *testing.T) {
	// Given
	c := canvas.New(10, 20)

	// Then
	assert.EqualValues(t, 10, c.Width)
	assert.EqualValues(t, 20, c.Height)
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			assert.EqualValues(t, 0, c.Pixels[i][j].Red)
			assert.EqualValues(t, 0, c.Pixels[i][j].Green)
			assert.EqualValues(t, 0, c.Pixels[i][j].Blue)
		}
	}
}

func TestCanvasWritePixel(t *testing.T) {
	// Given
	c := canvas.New(20, 10)
	r := colour.New(1, 0, 0)

	// When
	c.WritePixel(2, 3, r)

	// Then
	assert.EqualValues(t, r, c.Pixels[2][3])
}

func TestCanvasToPPMHeader(t *testing.T) {
	// Given
	c := canvas.New(80, 50)

	// When
	ppm := c.CanvasToPPM()

	// Then
	lines := strings.Split(string(ppm), "\n")
	assert.EqualValues(t, "P3", lines[0])
	assert.EqualValues(t, "80 50", lines[1])
	assert.EqualValues(t, "255", lines[2])
}

func TestCanvasToPPMBody(t *testing.T) {
	// Given
	c := canvas.New(5, 3)
	p1 := colour.New(1.5, 0, 0)
	p2 := colour.New(0, 0.5, 0)
	p3 := colour.New(-0.5, 0, 1)

	// When
	c.WritePixel(0, 0, p1)
	c.WritePixel(2, 1, p2)
	c.WritePixel(4, 2, p3)
	ppm := c.CanvasToPPM()

	// Then
	lines := strings.Split(string(ppm), "\n")
	a1 := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0"
	a2 := "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0"
	a3 := "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255"

	assert.EqualValues(t, a1, lines[3])
	assert.EqualValues(t, a2, lines[4])
	assert.EqualValues(t, a3, lines[5])
}
