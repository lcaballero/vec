package f32

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func AssertEqualMatrices(t *testing.T, a, b Matrix) {
	assert.True(t, EqualMatrices(a, b))
}

func TestMatrix_By_Ident_Equals_Original(t *testing.T) {
	a := Matrix{
		1.0, -5.0,  3.0,
		0.0, -2.0,  6.0,
		7.0,  2.0, -4.0,
	}
	b := Ident()
	AssertEqualMatrices(t, a, a.Mult(b))
	AssertEqualMatrices(t, a, b.Mult(a))
}

func TestMatrix_Mult_Simple(t *testing.T) {
	a := Matrix{
		 1.0, -5.0,  3.0,
		 0.0, -2.0,  6.0,
		 7.0,  2.0, -4.0,
	}
	b := Matrix{
		-8.0, 6.0,  1.0,
		 7.0, 0.0, -3.0,
		 2.0, 4.0,  5.0,
	}
	c := Mult(a, b)

	AssertEqualMatrices(t, c,
		Matrix{
			-37.0, 18.0, 31.0,
			-2.0, 24.0, 36.0,
			-50.0, 26.0, -19.0,
		},
	)
}

func TestMatrix_Transpose_Swaps_R_for_C(t *testing.T) {
	a := Matrix{
		1.0, 2.0, 3.0,
		4.0, 5.0, 6.0,
		7.0, 8.0, 9.0,
	}
	b := a.Transpose()
	AssertEqualMatrices(t, b,
		Matrix{
			1.0, 4.0, 7.0,
			2.0, 5.0, 8.0,
			3.0, 6.0, 9.0,
		},
	)
}

func TestMatrix_Transpose_Identity_of_3x3_is_Ident(t *testing.T) {
	a := Ident()
	b := a.Transpose()
	AssertEqualMatrices(t, a, b)
}
