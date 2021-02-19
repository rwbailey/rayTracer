package shape

import "github.com/rwbailey/ray/ray"

type Shape interface {
	Intersect(r ray.Ray) []*Intersection
}
