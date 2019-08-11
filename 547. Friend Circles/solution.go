func findCircleNum(M [][]int) int {
	ans, visited := 0, make([]bool, len(M))
	for i := range M {
		if !visited[i] {
			ans++
			visitFriends(M, i, visited)
		}
	}
	return ans
}
func visitFriends(M [][]int, i int, visited []bool) {
	if visited[i] == true {
		return
	}
	visited[i] = true
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 {
			visitFriends(M, j, visited)
		}
	}
}
