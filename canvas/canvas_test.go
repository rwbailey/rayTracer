package canvas_test

import (
	"testing"

	"github.com/rwbailey/ray/canvas"
	"github.com/rwbailey/ray/colour"
	"github.com/stretchr/testify/assert"
)

func TestCanvasCreation(t *testing.T) {
	// Given
	c := canvas.New(10, 20)

	// Then
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
