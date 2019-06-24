func romanToInt(s string) int {
	var ans int
	symbol := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := len(s) - 1; i >= 0; i-- {
		if i < len(s)-1 && symbol[s[i]] < symbol[s[i+1]] {
			ans -= symbol[s[i]]
		} else {
			ans += symbol[s[i]]
		}
	}
	return ans
}
