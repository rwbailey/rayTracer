package shape

import "github.com/rwbailey/ray/ray"

type Sphere struct{}

func NewSphere() *Sphere {
	return &Sphere{}
}

func (*Sphere) ConvertRayToObjectSpace(r ray.Ray) ray.Ray {
	return r
}
