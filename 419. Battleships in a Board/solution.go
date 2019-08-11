func countBattleships(board [][]byte) int {
	ans := 0
	for i, row := range board {
		for j, ele := range row {
			if ele != 'X' {
				continue
			}
			if (i == 0 || board[i-1][j] != 'X') && (j == 0 || row[j-1] != 'X') {
				ans++
			}
		}
	}
	return ans
}
