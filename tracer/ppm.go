package tracer

import "os"

type PPM string

func (p PPM) Save(path string) error {
	return os.WriteFile(path, []byte(p), 0644)
}
