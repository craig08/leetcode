func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 || grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1 {
		return -1
	}
	queue, step := [][]int{}, 0
	diff := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	// use dummy cell (-1, -1) to leverage the same logic for starting cell (0, 0)
	queue = append(queue, []int{-1, -1})
	for len(queue) > 0 {
		for i := len(queue); i > 0; i-- {
			cell := queue[0]
			queue = queue[1:]
			for _, d := range diff {
				if x, y := cell[0]+d[0], cell[1]+d[1]; x < 0 || y < 0 || x >= len(grid) || y >= len(grid) {
					continue
				} else if grid[x][y] != 0 {
					continue
				} else if x == len(grid)-1 && y == len(grid)-1 {
					return step + 1
				} else {
					grid[x][y] = step + 1
					queue = append(queue, []int{x, y})
				}
			}
		}
		step++
	}
	return -1
}
