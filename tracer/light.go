package tracer

type PointLight struct {
	Position  Tuple
	Intensity Colour
}

func NewPointLight(p Tuple, i Colour) *PointLight {
	return &PointLight{
		Position:  p,
		Intensity: i,
	}
}
