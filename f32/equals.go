package f32

func EqualMatrices(a, b Matrix) bool {
	an := len(a)
	bn := len(b)
	if an != bn {
		return false
	}
	for i := 0; i < an; i++ {
		if a[0] != b[0] {
			return false
		}
	}
	return true
}
