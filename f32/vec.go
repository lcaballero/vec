package f32

import (
	"math"
	"fmt"
	"errors"
)

var ErrFindingRadianAngleBetweenVectors = errors.New(
	`error finding radians between vectors
	where one or both vectors is a zero vector.`,
)

// Vec is a vector of float32 with dimensionality of 4.
type Vec [4]float32

// Zero is a Vec where x, y, z, w are all 0.0.
var Zero = Vec{0.0, 0.0, 0.0, 0.0}

// X provides the x component of the vector.
func (v Vec) X() float32 {
	return v[0]
}

// Y provides the y component of the vector.
func (v Vec) Y() float32 {
	return v[1]
}

// Z provides the z component of the vector.
func (v Vec) Z() float32 {
	return v[2]
}

// W provides the w component of the vector.
func (v Vec) W() float32 {
	return v[3]
}

// Len provides the dimensionality of the Vector which should always be 4.
func (v Vec) Dim() int {
	return len(v)
}

// Equals compares this vector against the parameter Vec component by
// component, and returns true if each component is equal (without any
// fuzz).  Otherwise Equals returns false
func (a Vec) Equals(b Vec) bool {
	return a[0] == b[0] &&
	a[1] == b[1] &&
	a[2] == b[2] &&
	a[3] == b[3] &&
	a.Dim() == b.Dim()
}

// SumOfSqrs returns the sum of squaring each component.
func (v Vec) SumOfSqrs() float64 {
	sumOfSquares := (v[0] * v[0]) + (v[1] * v[1]) + (v[2] * v[2]) + (v[3] * v[3])
	return float64(sumOfSquares)
}

// Mag returns the square root of the sum of squares which is
// the magnitude of the vector.
func (v Vec) Mag() float64 {
	return math.Sqrt(v.SumOfSqrs())
}

// Scale multiplies each component by the value s.
func (v Vec) Scale(s float32) Vec {
	v[0] *= s; v[1] *= s; v[2] *= s; v[3] *= s
	return v
}

// String produces a string in the form: <x,y,z,w>.
func (v Vec) String() string {
	return fmt.Sprintf("<%f, %f, %f, %f>", v.X(), v.Y(), v.Z(), v.W())
}

// Normalize produces a unit vector from v.  However, if the vector is the
// zero vector then it a zero vector is returned.  Normalizing is
// mathematically: (1 / magnitude) * v.  Which explains why Normalizing
// the Zero vector would cause a division by 0 and why the zero vector is
// returned. (May change to return an error in the future.)
// TODO: consider returning an error here with the vector, to handle
// TODO: normalizing a zero vector.
func (v Vec) Normalize() Vec {
	d := v.Mag()
	if d == 0.0 {
		return Vec{}
	}
	return v.Scale(float32(1.0 / d))
}

// Add returns a new vector whereby the two vectors produce a new vector
// from adding each corresponding component of the two vectors.
func (a Vec) Add(b Vec) Vec {
	return Add(a, b)
}

// Sub is a convenience for a.Add(b.Negate()).
func (a Vec) Sub(b Vec) Vec {
	return Sub(a, b)
}

// Negate returns a new vector that with the same magnitude just pointing
// in the opposite direction.
func (a Vec) Negate() Vec {
	return a.Scale(-1.0)
}

// Dot computes the dot-product between this vector and the provided vector.
func (a Vec) Dot(b Vec) float64 {
	return Dot(a, b)
}

// RadsBetween returns the angle (in radians) between the two vectors via
// the equation acos( a.b / (|a|*|b|) ).  If the magnitude of either a or b
// is 0 then an error is produced.  Otherwise the angle in rads is returned.
func (a Vec) RadsBetween(b Vec) (float64, error) {
	return RadsBetween(a, b)
}

// Cross produces the cross-product of the two vectors though they were
// only 3d (minus the w element).
func (a Vec) Cross(b Vec) Vec {
	return Cross(a, b)
}

