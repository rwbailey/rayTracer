package image

import "io/ioutil"

type PPM string

func (p PPM) Save(path string) error {
	return ioutil.WriteFile(path, []byte(p), 0644)
}
