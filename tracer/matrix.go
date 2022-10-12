package tracer

import (
	"math"
)

type Matrix [][]float64

func NewMatrix(s [][]float64) Matrix {
	return Matrix(s)
}

func ZeroMatrix(n int) Matrix {
	s := [][]float64{}
	for i := 0; i < n; i++ {
		s = append(s, make([]float64, n))
	}
	return NewMatrix(s)
}

func IdentityMatrix(n int) Matrix {
	z := ZeroMatrix(n)
	for i := 0; i < n; i++ {
		z[i][i] = 1
	}
	return z
}

func (m1 Matrix) Equals(m2 Matrix) bool {
	for row := 0; row < len(m1); row++ {
		for col := 0; col < len(m1[row]); col++ {
			if !FloatEquals(m1[row][col], m2[row][col]) {
				return false
			}
		}
	}
	return true
}

// We only need to multiply 4x4 matrices
func (a Matrix) MultiplyMatrix(b Matrix) Matrix {
	ab := ZeroMatrix(4)

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

func (m Matrix) MultiplyTuple(t Tuple) Tuple {
	mt := NewTuple(0, 0, 0, 0)

	mt.X = m[0][0]*t.X + m[0][1]*t.Y + m[0][2]*t.Z + m[0][3]*t.W
	mt.Y = m[1][0]*t.X + m[1][1]*t.Y + m[1][2]*t.Z + m[1][3]*t.W
	mt.Z = m[2][0]*t.X + m[2][1]*t.Y + m[2][2]*t.Z + m[2][3]*t.W
	mt.W = m[3][0]*t.X + m[3][1]*t.Y + m[3][2]*t.Z + m[3][3]*t.W

	return mt
}

func (m Matrix) Transpose() Matrix {
	d := len(m)
	t := ZeroMatrix(d)

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

func (m Matrix) Inverse() Matrix {
	if !m.IsInvertable() {
		panic("The Matrix cannot be inverted (Det == 0)")
	}
	d := len(m)
	inv := ZeroMatrix(d)

	for row := 0; row < d; row++ {
		for col := 0; col < d; col++ {
			c := m.Cofactor(row, col)
			inv[col][row] = c / m.Determinant()
		}
	}
	return inv
}

/**
 *	Transformations
 */

// Does the same as MultiplyTuple
func (m Matrix) Transform(t Tuple) Tuple {
	mt := NewTuple(0, 0, 0, 0)

	mt.X = m[0][0]*t.X + m[0][1]*t.Y + m[0][2]*t.Z + m[0][3]*t.W
	mt.Y = m[1][0]*t.X + m[1][1]*t.Y + m[1][2]*t.Z + m[1][3]*t.W
	mt.Z = m[2][0]*t.X + m[2][1]*t.Y + m[2][2]*t.Z + m[2][3]*t.W
	mt.W = m[3][0]*t.X + m[3][1]*t.Y + m[3][2]*t.Z + m[3][3]*t.W

	return mt
}

func TranslationMatrix(x, y, z float64) Matrix {
	t := IdentityMatrix(4)
	t[0][3] = x
	t[1][3] = y
	t[2][3] = z
	return t
}

func (m Matrix) Translate(x, y, z float64) Matrix {
	return TranslationMatrix(x, y, z).MultiplyMatrix(m)
}

func ScalingMatrix(x, y, z float64) Matrix {
	s := IdentityMatrix(4)
	s[0][0] = x
	s[1][1] = y
	s[2][2] = z
	return s
}

func (m Matrix) Scale(x, y, z float64) Matrix {
	return ScalingMatrix(x, y, z).MultiplyMatrix(m)
}

func RotationXMatrix(r float64) Matrix {
	a := IdentityMatrix(4)
	a[1][1] = math.Cos(r)
	a[1][2] = -math.Sin(r)
	a[2][1] = math.Sin(r)
	a[2][2] = math.Cos(r)
	return a
}

func (m Matrix) RotateX(r float64) Matrix {
	return RotationXMatrix(r).MultiplyMatrix(m)
}

func RotationYMatrix(r float64) Matrix {
	a := IdentityMatrix(4)
	a[0][0] = math.Cos(r)
	a[0][2] = math.Sin(r)
	a[2][0] = -math.Sin(r)
	a[2][2] = math.Cos(r)
	return a
}

func (m Matrix) RotateY(r float64) Matrix {
	return RotationYMatrix(r).MultiplyMatrix(m)
}

func RotationZMatrix(r float64) Matrix {
	a := IdentityMatrix(4)
	a[0][0] = math.Cos(r)
	a[0][1] = -math.Sin(r)
	a[1][0] = math.Sin(r)
	a[1][1] = math.Cos(r)
	return a
}

func (m Matrix) RotateZ(r float64) Matrix {
	return RotationZMatrix(r).MultiplyMatrix(m)
}

func ShearingMatrix(xy, xz, yx, yz, zx, zy float64) Matrix {
	a := IdentityMatrix(4)
	a[0][1] = xy
	a[0][2] = xz
	a[1][0] = yx
	a[1][2] = yz
	a[2][0] = zx
	a[2][1] = zy
	return a
}

func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	return ShearingMatrix(xy, xz, yx, yz, zx, zy).MultiplyMatrix(m)
}

func ViewTransform(from, to, up Tuple) Matrix {
	forward := to.Subtract(from).Normalise()
	upn := up.Normalise()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)
	orientation := NewMatrix([][]float64{
		{left.X, left.Y, left.Z, 0},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})
	return orientation.MultiplyMatrix(TranslationMatrix(-from.X, -from.Y, -from.Z))
}
