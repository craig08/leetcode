func search(nums []int, target int) int {
	for l, r := 0, len(nums); l < r; {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
