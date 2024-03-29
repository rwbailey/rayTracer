package tracer

import (
	"math"
	"sync"
)

type Camera struct {
	HSize       float64
	VSize       float64
	FieldOfView float64
	Transform   Matrix
	PixelSize   float64
	HalfHeight  float64
	HalfWidth   float64
}

func NewCamera(h, v, f float64) *Camera {
	c := &Camera{
		HSize:       h,
		VSize:       v,
		FieldOfView: f,
		Transform:   IdentityMatrix(4),
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

func (c *Camera) RayForPixel(px, py int) Ray {
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
	pixel := c.Transform.Inverse().MultiplyTuple(Point(worldX, worldY, -1))
	origin := c.Transform.Inverse().MultiplyTuple(Point(0, 0, 0))
	direction := pixel.Subtract(origin).Normalise()

	return NewRay(origin, direction)
}

func (c *Camera) Render(w *World) *Canvas {
	img := NewCanvas(int(c.HSize), int(c.VSize))

	concurrency := 4
	var wg sync.WaitGroup
	wg.Add(concurrency)

	go func() {
		defer wg.Done()
		for y := 0; y < int(c.VSize)/2; y++ {
			for x := 0; x < int(c.HSize)/2; x++ {
				ray := c.RayForPixel(x, y)
				colour := w.ColourAt(ray)
				img.WritePixel(x, y, colour)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for y := int(c.VSize) / 2; y < int(c.VSize); y++ {
			for x := 0; x < int(c.HSize)/2; x++ {
				ray := c.RayForPixel(x, y)
				colour := w.ColourAt(ray)
				img.WritePixel(x, y, colour)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for y := 0; y < int(c.VSize)/2; y++ {
			for x := int(c.HSize) / 2; x < int(c.HSize); x++ {
				ray := c.RayForPixel(x, y)
				colour := w.ColourAt(ray)
				img.WritePixel(x, y, colour)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for y := int(c.VSize) / 2; y < int(c.VSize); y++ {
			for x := int(c.HSize) / 2; x < int(c.HSize); x++ {
				ray := c.RayForPixel(x, y)
				colour := w.ColourAt(ray)
				img.WritePixel(x, y, colour)
			}
		}
	}()

	wg.Wait()

	return img
}
