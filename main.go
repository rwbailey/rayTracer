package main

import (
	"fmt"
	"math"

	"github.com/rwbailey/ray/canvas"
	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/matrix"
	p "github.com/rwbailey/ray/projectile"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
	t "github.com/rwbailey/ray/tuple"
)

var white colour.Colour

func main() {

	can := canvas.New(900, 600)
	white = colour.New(1, 1, 1)

	circle(can)

	err := can.CanvasToPPM().Save("image.ppm")
	if err != nil {
		fmt.Println(err)
	}
}

func circle(can *canvas.Canvas) {
	rayOrigin := tuple.Point(0, 0, -5)
	wallZ := float64(10)
	wallSize := float64(7)
	canvasPixels := 500
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	can.Width = canvasPixels
	can.Height = canvasPixels
	colour := colour.New(1, 0, 0)
	s := shape.NewSphere()

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := tuple.Point(worldX, worldY, wallZ)

			r := ray.New(rayOrigin, position.Subtract(rayOrigin).Normalise())
			xs := s.Intersect(r)

			if shape.Hit(xs) != nil {
				can.WritePixel(x, y, colour)
			}
		}
	}
}

func clock(can *canvas.Canvas) {
	twelve := tuple.Point(0, 200, 0)

	for i := 0; i < 12; i++ {
		can.WritePixel(int(twelve.X)+450, int(twelve.Y)+300, white)
		rotZ := matrix.Identity(4).RotateZ(math.Pi / 6)
		twelve = rotZ.Transform(twelve)
	}
}

func projectileMotion(can *canvas.Canvas) {
	env := p.Environment{
		Gravity: t.Vector(0, -0.1, 0),
		Wind:    t.Vector(-0.01, 0, 0),
	}

	proj := p.Projectile{
		Position: t.Point(0, 1, 0),
		Velocity: t.Vector(1, 1.8, 0).Normalise().Multiply(11.25),
	}

	for proj.Position.Y >= 0 {
		x := int(proj.Position.X)
		y := int(proj.Position.Y)

		if x > 899 || y > 549 {
			proj = p.Tick(env, proj)
			continue
		}

		can.Pixels[899-x][549-y] = white
		fmt.Println(proj.Position.Y)
		proj = p.Tick(env, proj)
	}
}
