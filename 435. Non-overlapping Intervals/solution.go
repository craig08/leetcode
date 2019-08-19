import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, max := 0, intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < max {
			ans++
		} else {
			max = intervals[i][1]
		}
	}
	return ans
}
