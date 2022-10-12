package tracer

type Shape interface {
	GetMaterial() *Material
	GetTransform() Matrix
	Intersect(r Ray) []*Intersection
	NormalAt(worldPoint Tuple) Tuple
	SetMaterial(m *Material)
	SetTransform(m Matrix)
}
