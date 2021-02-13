package main

import (
	"fmt"
	"io/ioutil"

	"github.com/rwbailey/ray/canvas"
	"github.com/rwbailey/ray/colour"
	p "github.com/rwbailey/ray/projectile"
	t "github.com/rwbailey/ray/tuple"
)

func main() {

	can := canvas.New(900, 550)
	white := colour.New(1, 1, 1)

	env := p.Environment{
		Gravity: t.Vector(0, -0.1, 0),
		Wind:    t.Vector(-0.01, 0, 0),
	}

	proj := p.Projectile{
		Position: t.Point(0, 1, 0),
		Velocity: t.Vector(1, 1.8, 0).Normalise().Multiply(11.25),
	}

	for proj.Position.Y >= 0 {
		can.Pixels[899-int(proj.Position.X)][549-int(proj.Position.Y)] = white
		fmt.Println(proj.Position.Y)
		proj = p.Tick(env, proj)
	}

	ppm := can.CanvasToPPM()
	err := ioutil.WriteFile("image.ppm", []byte(ppm), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
