package f32

const (
	e11 int = iota
	e12
	e13
	e21
	e22
	e23
	e31
	e32
	e33
)

func Mult(a, b Matrix) Matrix {
	a11, a12, a13 := a.Row1()
	a21, a22, a23 := a.Row2()
	a31, a32, a33 := a.Row3()

	b11, b12, b13 := b.Row1()
	b21, b22, b23 := b.Row2()
	b31, b32, b33 := b.Row3()

	return Matrix{
		a11*b11 + a12*b21 + a13*b31, a11*b12 + a12*b22 + a13*b32, a11*b13 + a12*b23 + a13*b33,
		a21*b11 + a22*b21 + a23*b31, a21*b12 + a22*b22 + a23*b32, a21*b13 + a22*b23 + a23*b33,
		a31*b11 + a32*b21 + a33*b31, a31*b12 + a32*b22 + a33*b32, a32*b13 + a32*b23 + a33*b33,
	}
}

func Rotate(a Matrix, rads float64) Matrix {
	panic(ErrNotImplementedYet)
}

