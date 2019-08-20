func coinChange(coins []int, amount int) int {
	count := make([]int, amount+1)
	for i := 1; i < len(count); i++ {
		count[i] = -1
	}
	for _, c := range coins {
		for n := 1; n < len(count); n++ {
			if n >= c && count[n-c] != -1 && (count[n] == -1 || count[n] > count[n-c]+1) {
				count[n] = count[n-c] + 1
			}
		}
	}
	return count[amount]
}
