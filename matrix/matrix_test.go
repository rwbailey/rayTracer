package matrix_test

import (
	"testing"

	"github.com/rwbailey/ray/matrix"
	"github.com/rwbailey/ray/tuple"
	"github.com/stretchr/testify/assert"
)

func TestMatrixConstructionAndInspection(t *testing.T) {
	// Given
	m := matrix.New([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})

	// Then
	assert.EqualValues(t, 4, m[0][3])
	assert.EqualValues(t, 5.5, m[1][0])
	assert.EqualValues(t, 7.5, m[1][2])
	assert.EqualValues(t, 11, m[2][2])
	assert.EqualValues(t, 13.5, m[3][0])
	assert.EqualValues(t, 15.5, m[3][2])
}

func TestMatrixComparison(t *testing.T) {
	// Given
	m1 := matrix.New([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	m2 := matrix.New([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	m3 := matrix.New([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.51},
	})
	m4 := matrix.New([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.500001},
	})

	// Then
	assert.True(t, m1.Equals(m2))
	assert.False(t, m1.Equals(m3))
	assert.True(t, m1.Equals(m4))
}

func TestMatrixMultiplicationByMatrix(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	b := matrix.New([][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})
	ab := matrix.New([][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})

	// Then
	assert.EqualValues(t, ab, a.MultiplyMatrix(b))
}

func TestMatrixMultiplicationByTuple(t *testing.T) {
	// Given
	m := matrix.New([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	a := tuple.New(1, 2, 3, 1)

	// Then
	assert.EqualValues(t, tuple.New(18, 24, 33, 1), m.MultiplyTuple(a))
}

func TestMatrixIdentityMultiplication(t *testing.T) {
	// Given
	i := matrix.Identity(4)
	m := matrix.New([][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	})

	// Then
	assert.EqualValues(t, m, i.MultiplyMatrix(m))
	assert.EqualValues(t, m, m.MultiplyMatrix(i))
}

func TestMatrixTranspose(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	})
	b := matrix.New([][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})
	z := matrix.Identity(3)

	// Then
	assert.EqualValues(t, b, a.Transpose())
	assert.EqualValues(t, z, z.Transpose())
}

func TestMatrixDeterminant2x2(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{1, 5},
		{-3, 2},
	})

	// Then
	assert.EqualValues(t, 17, a.Determinant())
}

func TestSubmatrixExtraction(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{1, 5, 3},
		{-3, 2, 7},
		{0, 6, -3},
	})

	s1 := matrix.New([][]float64{
		{-3, 2},
		{0, 6},
	})

	// Then
	assert.EqualValues(t, s1, a.Submatrix(0, 2))

	// Given
	b := matrix.New([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})

	s2 := matrix.New([][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})

	// Then
	assert.EqualValues(t, s2, b.Submatrix(2, 1))
}

func Test3x3MatrixMinor(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	b := a.Submatrix(1, 0)

	// Then
	assert.EqualValues(t, 25, b.Determinant())
	assert.EqualValues(t, 25, a.Minor(1, 0))
}

func TestMatrixCofactor(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	// Then
	assert.EqualValues(t, -12, a.Minor(0, 0))
	assert.EqualValues(t, -12, a.Cofactor(0, 0))
	assert.EqualValues(t, 25, a.Minor(1, 0))
	assert.EqualValues(t, -25, a.Cofactor(1, 0))
}

func TestMatrixDeterminant(t *testing.T) {
	// 3x3
	// Given
	a := matrix.New([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	// Then
	assert.EqualValues(t, 56, a.Cofactor(0, 0))
	assert.EqualValues(t, 12, a.Cofactor(0, 1))
	assert.EqualValues(t, -46, a.Cofactor(0, 2))
	assert.EqualValues(t, -196, a.Determinant())

	// 4x4
	// Given
	b := matrix.New([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})

	// Then
	assert.EqualValues(t, 690, b.Cofactor(0, 0))
	assert.EqualValues(t, 447, b.Cofactor(0, 1))
	assert.EqualValues(t, 210, b.Cofactor(0, 2))
	assert.EqualValues(t, 51, b.Cofactor(0, 3))
	assert.EqualValues(t, -4071, b.Determinant())
}

func TestMatrixInvertibility(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})
	b := matrix.New([][]float64{
		{-4, 2, -2, 3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})

	// Then
	assert.True(t, a.IsInvertable())
	assert.False(t, b.IsInvertable())
}

func TestMatrixInversion(t *testing.T) {
	// Given
	a := matrix.New([][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})
	b, _ := a.Inverse()
	B := matrix.New([][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	})

	// Then
	assert.EqualValues(t, 532, a.Determinant())
	assert.EqualValues(t, -160, a.Cofactor(2, 3))
	assert.EqualValues(t, -160.0/532, b[3][2])

	assert.EqualValues(t, 105, a.Cofactor(3, 2))
	assert.EqualValues(t, 105.0/532, b[2][3])
	assert.True(t, b.Equals(B))
}

func TestMoreInversions(t *testing.T) {
	// Given
	A := matrix.New([][]float64{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	})
	B := matrix.New([][]float64{
		{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308},
	})

	// Then
	iA, _ := A.Inverse()
	assert.True(t, iA.Equals(B))

	// Given
	C := matrix.New([][]float64{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	})
	D := matrix.New([][]float64{
		{-0.04074, -0.07778, 0.14444, -0.22222},
		{-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963},
		{0.17778, 0.06667, -0.26667, 0.33333},
	})

	// Then
	iC, _ := C.Inverse()
	assert.True(t, iC.Equals(D))
}

func TestMultiplyMatrixProductByInverse(t *testing.T) {
	// Given
	A := matrix.New([][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})
	B := matrix.New([][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})
	C := A.MultiplyMatrix(B)
	D := B.MultiplyMatrix(A)

	// Then
	iB, _ := B.Inverse()
	assert.True(t, C.MultiplyMatrix(iB).Equals(A))

	iA, _ := A.Inverse()
	assert.True(t, D.MultiplyMatrix(iA).Equals(B))
}

func TestMultiplyByTransformationMatrix(t *testing.T) {
	// Given
	trans := matrix.Identity(4).Translate(5, -3, 2)
	p := tuple.Point(-3, 4, 5)

	// Then
	assert.True(t, tuple.Point(2, 1, 7).Equals(trans.MultiplyTuple(p)))
}

func TestMultiplyByInverseOfTransformationMatrix(t *testing.T) {
	// Given
	trans := matrix.Identity(4).Translate(5, -3, 2)
	inv, _ := trans.Inverse()
	p := tuple.Point(-3, 4, 5)

	// Then
	assert.True(t, tuple.Point(-8, 7, 3).Equals(inv.MultiplyTuple(p)))
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	// Given
	trans := matrix.Identity(4).Translate(5, -3, 2)
	v := tuple.Vector(-3, 4, 5)

	// Then
	assert.True(t, v.Equals(trans.MultiplyTuple(v)))
}

func TestScalingMatrixAppliedToPoint(t *testing.T) {
	// Given
	scale := matrix.Identity(4).Scale(2, 3, 4)
	p := tuple.Point(-4, 6, 8)

	// Then
	assert.True(t, tuple.Point(-8, 18, 32).Equals(scale.MultiplyTuple(p)))
}

func TestScalingMatrixAppliedToVector(t *testing.T) {
	// Given
	scale := matrix.Identity(4).Scale(2, 3, 4)
	v := tuple.Vector(-4, 6, 8)

	// Then
	assert.True(t, tuple.Vector(-8, 18, 32).Equals(scale.MultiplyTuple(v)))
}

func TestMultiplyByInverseOfScalingMatrix(t *testing.T) {
	// Given
	scale := matrix.Identity(4).Scale(2, 3, 4)
	v := tuple.Vector(-4, 6, 8)
	inv, _ := scale.Inverse()

	// Then
	assert.True(t, tuple.Vector(-2, 2, 2).Equals(inv.MultiplyTuple(v)))
}

func TestReflectionIsScalingByANegativeValue(t *testing.T) {
	// Given
	scale := matrix.Identity(4).Scale(-1, 1, 1)
	v := tuple.Vector(2, 3, 4)

	// Then
	assert.True(t, tuple.Vector(-2, 3, 4).Equals(scale.MultiplyTuple(v)))
}
