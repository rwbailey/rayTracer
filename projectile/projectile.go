package projectile

import (
	"github.com/rwbailey/ray/tuple"
)

type Projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

type Environment struct {
	Gravity tuple.Tuple
	Wind    tuple.Tuple
}

func Tick(e Environment, p Projectile) Projectile {
	position := p.Position.Add(p.Velocity)
	velocity := p.Velocity.Add(e.Gravity).Add(e.Wind)
	return Projectile{position, velocity}
}
