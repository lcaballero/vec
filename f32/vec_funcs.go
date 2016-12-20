package f32
import "math"


// Dot computes the dot-product between this vector and the provided vector.
func Dot(a, b Vec) float64 {
	return float64((a[0] * b[0]) + (a[1] * b[1]) + (a[2] * b[2]))
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
	return math.Acos(dot / (am * bm)), nil
}

// Add returns a new vector whereby the two vectors produce a new vector
// from adding each corresponding component of the two vectors.
func Add(a, b Vec) Vec {
	return Vec{ a[0] + b[0], a[1] + b[1], a[2] + b[2] }
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
