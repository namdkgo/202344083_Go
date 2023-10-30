package mymath

func Myabs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Mypow(base int, exponent int) float64 {
	var num float64 = 1.0
	if exponent > 0 {
		b := float64(base)
		for i := 0; i < exponent; i++ {
			num *= b
		}
	} else if exponent < 0 {
		b := float64(base)
		for i := exponent; i < 0; i++ {
			num /= b
		}
	}
	return num
}
