func isPalindrome(x int) bool {
	o
	digits := 1
	if x < 0 {
		return false
	}
	for num := 10; num <= x; num, digits = num*10, digits+1 {
	}
	var nums []int
	for d := 0; d < digits/2; d, x = d+1, x/10 {
		nums = append(nums, x%10)
	}
	if digits&1 == 1 {
		x /= 10
	}
	for d := 0; d < digits/2; d, x = d+1, x/10 {
		if x%10 != nums[len(nums)-1] {
			return false
		}
		nums = nums[:len(nums)-1]
	}
	return true
}
