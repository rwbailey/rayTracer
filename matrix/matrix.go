package matrix

type Matrix [][]float64

func New(s [][]float64) Matrix {
	return Matrix(s)
}
