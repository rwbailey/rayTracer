package tracer

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r Ray) Transform(m Matrix) Ray {
	return Ray{
		Origin:    m.Transform(r.Origin),
		Direction: m.Transform(r.Direction),
	}
}
