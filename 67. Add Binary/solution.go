func addBinary(a string, b string) string {
	carry, ans := byte(0), ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; {
		sum := carry
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		ans = string('0'+sum) + ans
	}
	return ans
}
