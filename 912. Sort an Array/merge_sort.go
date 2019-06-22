func sortArray(nums []int) []int {
	mergesort(nums, 0, len(nums))
	return nums
}

func mergesort(nums []int, start, end int) {
	if start >= end-1 {
		return
	}
	mergesort(nums, start, (start+end)/2)
	mergesort(nums, (start+end)/2, end)
	merge(nums, start, end)
}

func merge(nums []int, start, end int) {
	arr, mid := make([]int, end-start), (start+end)/2
	for i, j, k := start, mid, 0; k < len(arr); k++ {
		if j == end || (i < mid && nums[i] < nums[j]) {
			arr[k] = nums[i]
			i++
		} else {
			arr[k] = nums[j]
			j++
		}
	}
	copy(nums[start:end], arr)
}
