package f32

import "math"

// Dot computes the dot-product between this vector and the provided vector.
func Dot(a, b Vec) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

// RadsBetween returns the angle (in radians) between the two vectors via
// the equation acos( a.b / (|a|*|b|) ).  If the magnitude of either a or b
// is 0 then an error is produced.  Otherwise the angle in rads is returned.
func RadsBetween(a, b Vec) (float64, error) {
	am := a.Mag()
	bm := b.Mag()
	if am == 0.0 || bm == 0.0 {
		return 0.0, ErrFindingRadianAngleBetweenVectors
	}
	dot := a.Dot(b)
	val := float64(dot) / (am * bm)
	return math.Acos(val), nil
}

// Add returns a new vector whereby the two vectors produce a new vector
// from adding each corresponding component of the two vectors.
func Add(a, b Vec) Vec {
	return Vec{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

// Sub is a convenience for a.Add(b.Negate()).
func Sub(a, b Vec) Vec {
	return Add(a, b.Negate())
}

// Cross produces the cross-product of the two vectors though they were
// only 3d (minus the w element).
func Cross(a, b Vec) Vec {
	return Vec{
		(a[1] * b[2]) - (a[2] * b[1]),
		(a[2] * b[0]) - (a[0] * b[2]),
		(a[0] * b[1]) - (a[1] * b[0]),
	}
}

// RotateMX creates a Matrix for rotating about the x-axis by the given
// angle provide as radians.
func RotateMX(rads float64) Matrix {
	return Matrix{
		1.0, 0.0, 0.0,
		0.0, cosf32(rads), sinf32(rads),
		0.0, -sinf32(rads), cosf32(rads),
	}
}

// RotateMY creates a Matrix for rotating about the y-axis by the given
// angle provide as radians.
func RotateMY(rads float64) Matrix {
	return Matrix{
		cosf32(rads), 0.0, -sinf32(rads),
		0.0, 1.0, 0.0,
		sinf32(rads), 0.0, cosf32(rads),
	}
}

// RotateMZ creates a Matrix for rotating about the z-axis by the given
// angle provide as radians.
func RotateMZ(rads float64) Matrix {
	return Matrix{
		cosf32(rads), sinf32(rads), 0.0,
		-sinf32(rads), cosf32(rads), 0.0,
		0.0, 0.0, 1.0,
	}
}

// RowMult multiplies the vector a by matrix b and returns the
// resulting row vector.
func RowMult(a Vec, b Matrix) Vec {
	return Vec{
		a[0]*b[e11] + a[1]*b[e21] + a[2]*b[e31],
		a[0]*b[e12] + a[1]*b[e22] + a[2]*b[e32],
		a[0]*b[e13] + a[1]*b[e23] + a[2]*b[e33],
	}
}

// ColMult multiplies the matrix b by vector a and return the
// resulting column vector.
func ColMult(b Matrix, a Vec) Vec {
	return Vec{
		a[0]*b[e11] + a[1]*b[e12] + a[2]*b[e13],
		a[0]*b[e21] + a[1]*b[e22] + a[2]*b[e23],
		a[0]*b[e31] + a[1]*b[e32] + a[2]*b[e33],
	}
}
