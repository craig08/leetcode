import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := IntSlice(nums[:k])
	heap.Init(&h)
	for i := k; i < len(nums); i++ {
		heap.Push(&h, nums[i])
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

type IntSlice []int

func (p IntSlice) Len() int            { return len(p) }
func (p IntSlice) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p IntSlice) Less(i, j int) bool  { return p[i] < p[j] }
func (p *IntSlice) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *IntSlice) Pop() interface{} {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}
