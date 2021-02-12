package canvas_test

import (
	"testing"

	. "github.com/rwbailey/ray/canvas"
	"github.com/stretchr/testify/assert"
)

func TestCanvasCreation(t *testing.T) {
	// Given
	c := New(10, 20)

	// Then
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			assert.EqualValues(t, 0, c.Pixels[i][j].Red)
			assert.EqualValues(t, 0, c.Pixels[i][j].Green)
			assert.EqualValues(t, 0, c.Pixels[i][j].Blue)
		}
	}
}
