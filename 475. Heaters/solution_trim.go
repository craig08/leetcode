import "sort"

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	radius := 0
	for hIdx := 0; hIdx < len(houses); {
		nearest := sort.Search(len(heaters), func(n int) bool {
			return heaters[n] >= houses[hIdx]
		})
		if nearest == 0 {
			radius = max(radius, abs(heaters[nearest]-houses[hIdx]))
		} else if nearest == len(heaters) {
			radius = max(radius, abs(heaters[nearest-1]-houses[hIdx]))
		} else {
			radius = max(radius, min(abs(heaters[nearest-1]-houses[hIdx]), abs(heaters[nearest]-houses[hIdx])))
		}
		for ; hIdx < len(houses); hIdx++ {
			if nearest > 0 && abs(heaters[nearest-1]-houses[hIdx]) <= radius {
				continue
			}
			if nearest < len(heaters) && abs(heaters[nearest]-houses[hIdx]) <= radius {
				continue
			}
			break
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
