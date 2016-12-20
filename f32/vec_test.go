package f32

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"math"
)

func TestVec_Cross(t *testing.T) {
	a := Vec{1.0, 3.0, 4.0}
	b := Vec{2.5, -5.0, 8.0}
	c := a.Cross(b)

	cx := (a.Y() * b.Z()) - (a.Z() * b.Y())
	cy := (a.Z() * b.X()) - (a.X() * b.Z())
	cz := (a.X() * b.Y()) - (a.Y() * b.X())

	assert.Equal(t, c.X(), float32(cx))
	assert.Equal(t, c.Y(), float32(cy))
	assert.Equal(t, c.Z(), float32(cz))
}

func TestVec_RadsBetween(t *testing.T) {
	theta := math.Pi / 2
	x := float32(math.Cos(theta))
	y := float32(math.Sin(theta))
	a := Vec{ x, 0.0, 0.0 }
	b := Vec{ 0.0, y, 0.0 }
	rads, err := a.RadsBetween(b)
	assert.Nil(t, err)
	t.Log(rads, theta)
	assert.Equal(t, rads, theta)
}

func TestVec_Dot(t *testing.T) {
	a := Vec{1.0, 2.0, 3.0}
	b := Vec{10.0, 20.0, 30.0}
	c := a.Dot(b)
	xs := a.X() * b.X()
	ys := a.Y() * b.Y()
	zs := a.Z() * b.Z()
	expectedDot := xs + ys + zs
	assert.Equal(t, c, expectedDot)
}

func TestVec_Sub(t *testing.T) {
	a := Vec{1.0, 2.0, 3.0}
	b := Vec{0.0, 1.0, 2.0}
	c := a.Sub(b)
	expected := float32(1.0)
	assert.Equal(t, c.X(), expected)
	assert.Equal(t, c.Y(), expected)
	assert.Equal(t, c.Z(), expected)
}

func TestVec_Negate(t *testing.T) {
	a := Vec{1.0, 2.0, 3.0}
	b := a.Negate()
	assert.Equal(t, a.X(), -b.X())
	assert.Equal(t, a.Y(), -b.Y())
	assert.Equal(t, a.Z(), -b.Z())
}

func TestVec_Add(t *testing.T) {
	a := Vec{1.0, 11.0, 222.0}
	b := Vec{-21.0, -42.0, 99.72}
	c := a.Add(b)
	assert.Equal(t, c.X(), a.X() + b.X())
	assert.Equal(t, c.Y(), a.Y() + b.Y())
	assert.Equal(t, c.Z(), a.Z() + b.Z())
}

func TestVec_Normalized_Vec_Has_Magnitude_of_One(t *testing.T) {
	a := Vec{1.0, 2.0, 3.0}
	n := a.Normalize()
	m := n.Mag()
	assert.InDelta(t, m, 1.0, 0.00001)
}

func TestVec_Scaling_Equals_Vector_With_Opposite_Direction(t *testing.T) {
	a := Vec{1.0, 2.0, 3.0}
	b := a.Scale(-1.0)

	c := Vec{-1.0, -2.0, -3.0}
	assert.True(t, b.Equals(c))

	// Shows: a is unchanged since it equals the reverse reversed.
	assert.True(t, a.Equals(c.Scale(-1.0)))
}

func TestVec_SumOfSqrs_For_1_2_3_4_Vec_Is_Correct(t *testing.T) {
	a := Vec{1.0, 2.0, 3.0}
	assert.Equal(t, a.SumOfSqrs(), 1.0 + 4.0 + 9.0)
}

func Test_Magnitude_Of_1_1_1_1_Vec_Is_2(t *testing.T) {
	a := Vec{1.0, 1.0, 1.0}
	mag := a.Mag()
	assert.Equal(t, mag, math.Sqrt(3.0))
}

func Test_Magnitude_Of_3_4_Vec_Is_5(t *testing.T) {
	a := Vec{3.0, 4.0, 0.0}
	mag := a.Mag()
	assert.Equal(t, mag, 5.0)
}

func Test_Each_Component_Getter_Produces_Right_Value(t *testing.T) {
	a := Vec{1.1, 2.2, 3.3}
	assert.Equal(t, a.X(), a[0])
	assert.Equal(t, a.Y(), a[1])
	assert.Equal(t, a.Z(), a[2])
}

func Test_A_Does_Not_Equal_B_With_Different_Components(t *testing.T) {
	a := Vec{1.1, 2.2, 3.3}
	b := Vec{1.1, 2.2, 33.0}
	assert.False(t, a.Equals(b))
}

func Test_A_Equals_B_With_Same_Components(t *testing.T) {
	a := Vec{1.1, 2.2, 3.3}
	b := Vec{1.1, 2.2, 3.3}
	assert.True(t, a.Equals(b))
}

func Test_Zero_Equals_New_Default_Vec(t *testing.T) {
	assert.True(t, Zero.Equals(Vec{}))
}

func Test_Default_Vec_Has_Dimension_Of_4(t *testing.T) {
	v := Vec{}
	assert.Equal(t, len(v), 3)
	assert.Equal(t, len(v), v.Dim())
}

func Test_Zero_Vec_XYZW_Are_Zero(t *testing.T) {
	zero := float32(0.0)
	assert.Equal(t, Zero[0], zero)
	assert.Equal(t, Zero[1], zero)
	assert.Equal(t, Zero[2], zero)
}

func Test_New_Vec_Is_Zero_Vec(t *testing.T) {
	v := Vec{}
	zero := float32(0.0)
	assert.Equal(t, v[0], zero)
	assert.Equal(t, v[1], zero)
	assert.Equal(t, v[2], zero)
}

func Test_Vec_String_Is_Vector_Component_String(t *testing.T) {
	v := Vec{1.22, 3.44, 5.66}
	s := v.String()
	assert.Contains(t, s, "<")
	assert.Contains(t, s, ">")
	assert.Contains(t, s, "1.22")
	assert.Contains(t, s, "3.44")
	assert.Contains(t, s, "5.66")
}