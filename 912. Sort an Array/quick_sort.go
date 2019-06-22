func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums))
	return nums
}

func quicksort(nums []int, start, end int) {
	if start >= end-1 {
		return
	}
	pivot := start + rand.Intn(end-start)
	nums[pivot], nums[end-1] = nums[end-1], nums[pivot]
	pos := partition(nums, start, end)
	quicksort(nums, start, pos)
	quicksort(nums, pos+1, end)
}

func partition(nums []int, start, end int) int {
	j := start
	for i := start; i < end-1; i++ {
		if nums[i] < nums[end-1] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[end-1], nums[j] = nums[j], nums[end-1]
	return j
}
