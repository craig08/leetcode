import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	radius := 0
	for hIdx := 0; hIdx < len(houses); hIdx++ {
		nearest := sort.Search(len(heaters), func(n int) bool {
			return heaters[n] >= houses[hIdx]
		})
		if nearest == 0 {
			radius = max(radius, heaters[nearest]-houses[hIdx])
		} else if nearest == len(heaters) {
			radius = max(radius, houses[hIdx]-heaters[nearest-1])
		} else {
			radius = max(radius, min(houses[hIdx]-heaters[nearest-1], heaters[nearest]-houses[hIdx]))
		}
	}
	return radius
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
