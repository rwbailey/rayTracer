package camera

import (
	"math"

	"github.com/rwbailey/ray/matrix"
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
