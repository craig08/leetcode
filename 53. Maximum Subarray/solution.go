func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, ans, max := 0, 0, nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		sum += num
		if sum > ans {
			ans = sum
		} else if sum < 0 {
			sum = 0
		}
	}
	if max < 0 {
		return max
	}
	return ans
}
