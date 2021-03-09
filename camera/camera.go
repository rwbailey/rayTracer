package camera

import "github.com/rwbailey/ray/matrix"

type Camera struct {
	HSize       float64
	VSize       float64
	FieldOfView float64
	Transform   matrix.Matrix
}

func New(h, v, f float64) *Camera {
	return &Camera{
		HSize:       h,
		VSize:       v,
		FieldOfView: f,
		Transform:   matrix.Identity(4),
	}
}
