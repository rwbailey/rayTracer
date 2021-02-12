package canvas

import "github.com/rwbailey/ray/colour"

type Canvas struct {
	Width  int
	Height int
	Pixels [][]colour.Colour
}

func New(w, h int) *Canvas {
	pixels := make([][]colour.Colour, w)

	for i := 0; i < w; i++ {
		pixels[i] = make([]colour.Colour, h)
		for j := 0; j < h; j++ {
			pixels[i][j] = colour.New(0, 0, 0)
		}
	}
	return &Canvas{
		Width:  w,
		Height: h,
		Pixels: pixels,
	}
}

func (c *Canvas) WritePixel(x, y int, colour colour.Colour) {
	c.Pixels[x][y] = colour
}
