func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for n := 2; n*n <= num; n++ {
		if num%n == 0 {
			sum += n
			if n*n != num {
				sum += num / n
			}
		}
	}
	return num == sum
}
