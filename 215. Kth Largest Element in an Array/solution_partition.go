import "math/rand"

func findKthLargest(nums []int, k int) int {
	for start, end := 0, len(nums); ; {
		idx := partition(nums, start, end)
		if idx == k-1 {
			break
		} else if idx > k-1 {
			end = idx
		} else {
			start = idx + 1
		}
	}
	return nums[k-1]
}
func partition(nums []int, start, end int) int {
	// randomly choose pivot to avoid worst cases
	pivot := start + rand.Intn(end-start)
	nums[end-1], nums[pivot] = nums[pivot], nums[end-1]
	// numbers with index less than idx will be greater than pivot
	idx := start
	for i := start; i < end-1; i++ {
		if nums[i] > nums[end-1] {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	nums[end-1], nums[idx] = nums[idx], nums[end-1]
	return idx
}

