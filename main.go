package main

import (
	"fmt"
	"math"

	"github.com/rwbailey/ray/camera"
	"github.com/rwbailey/ray/canvas"
	"github.com/rwbailey/ray/colour"
	"github.com/rwbailey/ray/light"
	"github.com/rwbailey/ray/material"
	"github.com/rwbailey/ray/matrix"
	p "github.com/rwbailey/ray/projectile"
	"github.com/rwbailey/ray/ray"
	"github.com/rwbailey/ray/shape"
	"github.com/rwbailey/ray/tuple"
	t "github.com/rwbailey/ray/tuple"
	"github.com/rwbailey/ray/world"
)

var white colour.Colour

func main() {

	can := canvas.New(900, 600)
	white = colour.White

	can = scene()

	err := can.CanvasToPPM().Save("image.ppm")
	if err != nil {
		fmt.Println(err)
	}
}

func scene() *canvas.Canvas {
	floor := shape.NewSphere()
	floor.Transform = matrix.Scaling(10, 0.01, 10)
	floor.Material = material.New()
	floor.Material.Colour = colour.New(1, 0.9, 0.9)
	floor.Material.Specular = 0

	leftWall := shape.NewSphere()
	leftWall.Transform = matrix.Translation(0, 0, 5).MultiplyMatrix(matrix.RotationY(-math.Pi / 4)).MultiplyMatrix(matrix.RotationX(math.Pi / 2)).MultiplyMatrix(matrix.Scaling(10, 0.01, 10))
	leftWall.Material = floor.Material

	rightWall := shape.NewSphere()
	rightWall.Transform = matrix.Translation(0, 0, 5).MultiplyMatrix(matrix.RotationY(math.Pi / 4)).MultiplyMatrix(matrix.RotationX(math.Pi / 2)).MultiplyMatrix(matrix.Scaling(10, 0.01, 10))
	rightWall.Material = floor.Material

	middle := shape.NewSphere()
	middle.Transform = matrix.Translation(-0.5, 1, 0.5)
	middle.Material = material.New()
	middle.Material.Colour = colour.New(0.1, 0, 1)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := shape.NewSphere()
	right.Transform = matrix.Translation(1.5, 0.5, -0.5).MultiplyMatrix(matrix.Scaling(0.5, 0.5, 0.5))
	right.Material = material.New()
	right.Material.Colour = colour.New(0.5, 1, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := shape.NewSphere()
	left.Transform = matrix.Translation(-1.5, 0.33, -0.75).MultiplyMatrix(matrix.Scaling(0.33, 0.33, 0.33))
	left.Material = material.New()
	left.Material.Colour = colour.New(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	w := world.New()
	w.Light = light.NewPointLight(tuple.Point(-10, 10, -10), colour.White)
	w.AddObjects(floor, leftWall, rightWall, middle, right, left)

	c := camera.New(800, 500, math.Pi/3)
	c.Transform = matrix.ViewTransform(tuple.Point(0, 1.5, -5), tuple.Point(0, 1, 0), tuple.Point(0, 1, 0))

	img := c.Render(w)

	return img
}

func circle(can *canvas.Canvas) {
	rayOrigin := tuple.Point(0, 0, -5)
	wallZ := float64(10)
	wallSize := float64(7)
	canvasPixels := 500
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	lightPosition := tuple.Point(-10, 10, -10)
	lightColour := colour.New(1, 1, 1)
	pointLight := light.NewPointLight(lightPosition, lightColour)

	can.Width = canvasPixels
	can.Height = canvasPixels
	s := shape.NewSphere()
	s.Material.Colour = colour.New(0, 0.2, 1)

	T := matrix.Identity(4).Scale(1, 0.5, 1)

	s.Transform = T

	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := tuple.Point(worldX, worldY, wallZ)

			r := ray.New(rayOrigin, position.Subtract(rayOrigin).Normalise())
			r.Direction = r.Direction.Normalise()
			xs := s.Intersect(r)

			if hit := shape.Hit(xs); hit != nil {
				point := r.Position(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := r.Direction.Negate()
				c := s.Material.Lighting(pointLight, point, eye, normal, false)
				can.WritePixel(x, y, c)
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
