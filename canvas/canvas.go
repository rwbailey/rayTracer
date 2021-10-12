package canvas

import (
	"fmt"
	"math"

	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/image"
)

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

func (c *Canvas) PixelAt(x, y int) colour.Colour {
	return c.Pixels[x][y]
}

// TODO: Refactor
func (c *Canvas) CanvasToPPM() image.PPM {
	ppm := []byte("P3\n" + fmt.Sprint(c.Width) + " " + fmt.Sprint(c.Height) + "\n255\n")

	for y := 0; y < c.Height; y++ {
		n := 0
		for x := 0; x < c.Width; x++ {

			red := c.Pixels[x][y].Red
			green := c.Pixels[x][y].Green
			blue := c.Pixels[x][y].Blue

			if red < 0 {
				red = 0
			}
			if green < 0 {
				green = 0
			}
			if blue < 0 {
				blue = 0
			}
			if red > 1 {
				red = 1
			}
			if green > 1 {
				green = 1
			}
			if blue > 1 {
				blue = 1
			}

			r := int(math.Round(red * 255))
			g := int(math.Round(green * 255))
			b := int(math.Round(blue * 255))

			ppm = append(ppm, []byte(fmt.Sprint(r)+" ")...)
			n += 4
			if n >= 67 {
				ppm = ppm[:len(ppm)-1]
				ppm = append(ppm, byte('\n'))
				n = 0
			}
			ppm = append(ppm, []byte(fmt.Sprint(g)+" ")...)
			n += 4
			if n >= 67 {
				ppm = ppm[:len(ppm)-1]
				ppm = append(ppm, byte('\n'))
				n = 0
			}
			ppm = append(ppm, []byte(fmt.Sprint(b)+" ")...)
			n += 4
			if n >= 67 {
				ppm = ppm[:len(ppm)-1]
				ppm = append(ppm, byte('\n'))
				n = 0
			}
		}
		n = 0
		ppm = ppm[:len(ppm)-1]
		ppm = append(ppm, byte('\n'))
	}
	return image.PPM(string(ppm))
}
