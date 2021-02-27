package material

import "github.com/rwbailey/ray/colour"

type Material struct {
	Colour    colour.Colour
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func New() *Material {
	return &Material{
		Colour:    colour.New(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}
