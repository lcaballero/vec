package f32
import "errors"

var ErrNotImplementedYet = errors.New("not implemented")

type Matrix [16]float32

func Ident() Matrix {
	return Matrix{
		1.0, 0.0, 0.0, 0.0,
		0.0, 1.0, 0.0, 0.0,
		0.0, 0.0, 1.0, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}
}

func (m Matrix) Transpose() Matrix {
	panic(ErrNotImplementedYet)
}

func (m Matrix) Scalar(a float32) Matrix {
	panic(ErrNotImplementedYet)
}

func (m Matrix) Scale(a float32) Matrix {
	panic(ErrNotImplementedYet)
}

func (a Matrix) Mult(b Matrix) Matrix {
	panic(ErrNotImplementedYet)
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
