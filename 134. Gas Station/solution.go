func canCompleteCircuit(gas []int, cost []int) int {
	var start, sum int
	for i, n := range gas {
		sum += n - cost[i]
		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}
	if start >= len(gas) {
		return -1
	}
	for i := 0; i < start; i++ {
		sum += gas[i] - cost[i]
		if sum < 0 {
			return -1
		}
	}
	return start
}

// gas = [1, 2, 8]
// cost = [2, 3, 1]
