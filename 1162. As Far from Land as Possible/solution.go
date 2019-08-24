func maxDistance(grid [][]int) int {
	// d is the grid of distance from water cells to nearest land cells
	q, d, ok := [][]int{}, make([][]int, len(grid)), false
	for i := range d {
		d[i] = make([]int, len(grid[0]))
	}
	// add all land cells to BFS queue
	// there will be an exception if the grid is full of lands
	// use ok variable to identify this situation
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			} else {
				ok = true
			}
		}
	}
	// return earlier if this is special case
	if !ok {
		return -1
	}
	step := 0
	for len(q) > 0 {
		step++
		for l := len(q); l > 0; l-- {
			n := q[0]
			q = q[1:]
			for _, diff := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				if i, j := n[0]+diff[0], n[1]+diff[1]; 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) {
					if grid[i][j] != 0 {
						continue
					}
					// only record distance if the water cell has not been touched
					// because of the BFS method, once the distance is written, it is the nearest
					if d[i][j] == 0 {
						d[i][j] = step
						q = append(q, []int{i, j})
					}
				}
			}
		}
	}
	return step - 1
}
