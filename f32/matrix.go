package f32
import "errors"

var ErrNotImplementedYet = errors.New("not implemented")

type Matrix [9]float32

func Ident() Matrix {
	return Matrix{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
	}
}

func (m Matrix) Row1() (float32, float32, float32) {
	return m[e11], m[e12], m[e13]
}

func (m Matrix) Row2() (float32, float32, float32) {
	return m[e21], m[e22], m[e23]
}

func (m Matrix) Row3() (float32, float32, float32) {
	return m[e31], m[e32], m[e33]
}

func (m Matrix) Transpose() Matrix {
	return Matrix{
		m[e11], m[e21], m[e31],
		m[e12], m[e22], m[e32],
		m[e13], m[e23], m[e33],
	}
}

func (m Matrix) Scalar(a float32) Matrix {
	return Matrix{
		a*m[e11], a*m[e12], a*m[e13],

	}
}

func (m Matrix) Scale(a float32) Matrix {
	panic(ErrNotImplementedYet)
}

func (a Matrix) Mult(b Matrix) Matrix {
	return Mult(a, b)
}

func (a Matrix) Rotate(rads float64) Matrix {
	panic(ErrNotImplementedYet)
}

func (a Matrix) ReflectX() Matrix {
	panic(ErrNotImplementedYet)
}

func (a Matrix) ReflectY() Matrix {
	panic(ErrNotImplementedYet)
}

func (a Matrix) OrthoProjection() Matrix {
	panic(ErrNotImplementedYet)
}
