func addNegabinary(arr1 []int, arr2 []int) []int {
	carry, ans := 0, []int{}
	for i, j := len(arr1)-1, len(arr2)-1; i >= 0 || j >= 0 || carry != 0; {
		if i >= 0 {
			carry += arr1[i]
			i--
		}
		if j >= 0 {
			carry += arr2[j]
			j--
		}
		ans = append([]int{carry & 1}, ans...)
		carry = -(carry >> 1)
	}
	for idx, num := range ans {
		if num != 0 {
			return ans[idx:]
		}
	}
	return []int{0}
}
