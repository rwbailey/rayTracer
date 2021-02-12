package main

import (
	"fmt"

	p "github.com/rwbailey/ray/projectile"
	t "github.com/rwbailey/ray/tuple"
)

func main() {
	env := p.Environment{
		Gravity: t.Vector(0, -0.1, 0),
		Wind:    t.Vector(-0.01, 0, 0),
	}

	proj := p.Projectile{
		Position: t.Point(0, 1, 0),
		Velocity: t.Vector(10, 10, 0),
	}

	for proj.Position.Y >= 0 {
		fmt.Println(proj.Position.X, proj.Position.Y)
		proj = p.Tick(env, proj)
	}
}
