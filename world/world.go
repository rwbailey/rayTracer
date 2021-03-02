package world

import (
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/shape"
)

type World struct {
	Objects []*shape.Shape
	Light   *light.PointLight
}

func New() *World {
	return &World{
		Objects: []*shape.Shape{},
		Light:   nil,
	}
}
