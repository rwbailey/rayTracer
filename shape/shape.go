package shape

import "github.com/rwbailey/ray/ray"

type Shape interface {
	ConvertRayToObjectSpace(r ray.Ray) ray.Ray
}
