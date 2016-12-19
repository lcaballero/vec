package f32

func (m Matrix) E11() float32 {
	return m[0]
}
func (m Matrix) E12() float32 {
	return m[1]
}
func (m Matrix) E13() float32 {
	return m[2]
}
func (m Matrix) E14() float32 {
	return m[3]
}

func (m Matrix) E21() float32 {
	return m[4]
}
func (m Matrix) E22() float32 {
	return m[5]
}
func (m Matrix) E23() float32 {
	return m[6]
}
func (m Matrix) E24() float32 {
	return m[7]
}

func (m Matrix) E31() float32 {
	return m[8]
}
func (m Matrix) E32() float32 {
	return m[9]
}
func (m Matrix) E33() float32 {
	return m[10]
}
func (m Matrix) E34() float32 {
	return m[11]
}

func (m Matrix) E41() float32 {
	return m[12]
}
func (m Matrix) E42() float32 {
	return m[13]
}
func (m Matrix) E43() float32 {
	return m[14]
}
func (m Matrix) E44() float32 {
	return m[15]
}
