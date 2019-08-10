func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	// for sub-arrays longer than 4 elements, we can find all different
	// indexes representing first number, first quartile, median, third quartile,
	// and the last number. Eliminate by a half for each loop.
	for l < r-4 {
		m := (l + r) / 2
		leftQ, rightQ := (l+m)/2, (m+r)/2
		if !(nums[l] < nums[leftQ] && nums[leftQ] < nums[m]) {
			r = m
		} else if !(nums[r] < nums[rightQ] && nums[rightQ] < nums[m]) {
			l = m
		} else {
			l, r = leftQ, rightQ
		}
	}
	// if the length is less than 5, the maximum in the sub-array is the answer
	for i := l; i <= r; i++ {
		if (i == l || nums[i-1] < nums[i]) && (i == r || nums[i] > nums[i+1]) {
			return i
		}
	}
	return 0
}
