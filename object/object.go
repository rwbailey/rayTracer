package object

type Object struct{}

type Sphere Object

func NewSphere() *Sphere {
	return &Sphere{}
}
