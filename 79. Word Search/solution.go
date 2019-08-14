func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			if find(board, word, used, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func find(board [][]byte, word string, used [][]bool, i, j, start int) bool {
	if start >= len(word) {
		return true
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return false
	}
	if used[i][j] || board[i][j] != word[start] {
		return false
	}
	used[i][j] = true
	diff := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range diff {
		if find(board, word, used, i+d[0], j+d[1], start+1) {
			return true
		}
	}
	used[i][j] = false
	return false
}
