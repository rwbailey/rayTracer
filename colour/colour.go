package colour

type colour struct {
	Red   float64
	Green float64
	Blue  float64
}

func New(r, g, b float64) colour {
	return colour{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}
