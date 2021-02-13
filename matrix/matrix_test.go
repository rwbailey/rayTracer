package matrix_test

import (
	"testing"

	"github.com/rwbailey/ray/matrix"
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
