func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / myPow(x, -n)
	}
	// check n's parity
	num := myPow(x, n>>1)
	if n&1 == 1 {
		return num * num * x
	}
	return num * num
}
