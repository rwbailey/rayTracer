package matrix

import (
	"errors"

	"github.com/rwbailey/ray/helpers"
	"github.com/rwbailey/ray/tuple"
)

type Matrix [][]float64

func New(s [][]float64) Matrix {
	return Matrix(s)
}

func Zero(n int) Matrix {
	s := [][]float64{}
	for i := 0; i < n; i++ {
		s = append(s, make([]float64, n))
	}
	return New(s)
}

func Identity(n int) Matrix {
	z := Zero(n)
	for i := 0; i < n; i++ {
		z[i][i] = 1
	}
	return z
}

func (m1 Matrix) Equals(m2 Matrix) bool {
	for row := 0; row < len(m1); row++ {
		for col := 0; col < len(m1[row]); col++ {
			if !helpers.FloatEquals(m1[row][col], m2[row][col]) {
				return false
			}
		}
	}
	return true
}

// We only need to multiply 4x4 matrices
func (a Matrix) MultiplyMatrix(b Matrix) Matrix {
	ab := Zero(4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			ab[row][col] = a[row][0]*b[0][col] +
				a[row][1]*b[1][col] +
				a[row][2]*b[2][col] +
				a[row][3]*b[3][col]
		}
	}

	return ab
}

func (m Matrix) MultiplyTuple(t tuple.Tuple) tuple.Tuple {
	mt := tuple.New(0, 0, 0, 0)

	mt.X = m[0][0]*t.X + m[0][1]*t.Y + m[0][2]*t.Z + m[0][3]*t.W
	mt.Y = m[1][0]*t.X + m[1][1]*t.Y + m[1][2]*t.Z + m[1][3]*t.W
	mt.Z = m[2][0]*t.X + m[2][1]*t.Y + m[2][2]*t.Z + m[2][3]*t.W
	mt.W = m[3][0]*t.X + m[3][1]*t.Y + m[3][2]*t.Z + m[3][3]*t.W

	return mt
}

func (m Matrix) Transpose() Matrix {
	d := len(m)
	t := Zero(d)

	for row := 0; row < d; row++ {
		for col := 0; col < d; col++ {
			t[row][col] = m[col][row]
		}
	}
	return t
}

func (m Matrix) Determinant() float64 {
	var det float64
	d := len(m)
	if d == 2 {
		det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	} else {
		for col := 0; col < d; col++ {
			det += m[0][col] * m.Cofactor(0, col)
		}
	}
	return det
}

func (m Matrix) Submatrix(r, c int) Matrix {
	d := len(m)
	s := make([][]float64, 0, d-1)

	for row := 0; row < d; row++ {
		x := make([]float64, 0, d-1)
		for col := 0; col < d; col++ {
			if col == c || row == r {
				continue
			}
			x = append(x, m[row][col])
		}
		if len(x) != 0 {
			s = append(s, x)
		}
	}
	return Matrix(s)
}

// 3x3 only
func (a Matrix) Minor(r, c int) float64 {
	return a.Submatrix(r, c).Determinant()
}

func (a Matrix) Cofactor(r, c int) float64 {
	min := a.Minor(r, c)
	if (r+c)%2 != 0 {
		return -min
	}
	return min
}

func (m Matrix) IsInvertable() bool {
	return m.Determinant() != 0
}

func (m Matrix) Inverse() (Matrix, error) {
	if !m.IsInvertable() {
		return nil, errors.New("The Matrix cannot be inverted (Det == 0)")
	}
	d := len(m)
	inv := Zero(d)

	for row := 0; row < d; row++ {
		for col := 0; col < d; col++ {
			c := m.Cofactor(row, col)
			inv[col][row] = c / m.Determinant()
		}
	}
	return inv, nil
}

func (m Matrix) Translate(x, y, z float64) Matrix {
	t := Identity(4)
	t[0][3] = x
	t[1][3] = y
	t[2][3] = z
	return t.MultiplyMatrix(m)
}
