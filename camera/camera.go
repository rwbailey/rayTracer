package camera

import (
	"math"

	"github.com/rwbailey/ray/canvas"
	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/tuple"
	"github.com/rwbailey/ray/world"
)

type Camera struct {
	HSize       float64
	VSize       float64
	FieldOfView float64
	Transform   matrix.Matrix
	PixelSize   float64
	HalfHeight  float64
	HalfWidth   float64
}

func New(h, v, f float64) *Camera {
	c := &Camera{
		HSize:       h,
		VSize:       v,
		FieldOfView: f,
		Transform:   matrix.Identity(4),
	}
	c.pixelSize(h, v, f)

	return c
}

func (c *Camera) pixelSize(h, v, f float64) {
	halfView := math.Tan(f / 2)
	aspect := h / v

	if aspect >= 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = (c.HalfWidth * 2.0) / h
}

func (c *Camera) RayForPixel(px, py int) ray.Ray {
	// The offset from the edge of te canvas to the pixel's centre
	xOffset := (float64(px) + 0.5) * c.PixelSize
	yOffset := (float64(py) + 0.5) * c.PixelSize

	// The transformed coordinates of the pixel in world space
	// (The camera looks towards -z, so +x is to the LEFT)
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	// Using the camera matrix, transform the canvas point and the origin,
	// and then compute the ray's direction vector.
	// (The canvas is at z=-1)
	pixel := c.Transform.Inverse().MultiplyTuple(tuple.Point(worldX, worldY, -1))
	origin := c.Transform.Inverse().MultiplyTuple(tuple.Point(0, 0, 0))
	direction := pixel.Subtract(origin).Normalise()

	return ray.New(origin, direction)
}

func (c *Camera) Render(w *world.World) *canvas.Canvas {
	img := canvas.New(int(c.HSize), int(c.VSize))

	for y := 0; y < int(c.VSize); y++ {
		for x := 0; x < int(c.HSize); x++ {
			ray := c.RayForPixel(x, y)
			colour := w.ColourAt(ray)
			img.WritePixel(x, y, colour)
		}
	}
	return img
}
