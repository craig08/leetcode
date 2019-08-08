import "sort"

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for first, num1 := range nums {
		if first > 0 && num1 == nums[first-1] {
			continue
		}
		for second, third := first+1, len(nums)-1; second < third; {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if sum := nums[second] + nums[third]; sum == -num1 {
				row := []int{num1, nums[second], nums[third]}
				ans = append(ans, row)
				second++
			} else if sum < -num1 {
				second++
			} else {
				third--
			}
		}
	}
	return ans
}
