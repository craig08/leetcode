func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	visitRooms(0, visited, rooms)
	return len(visited) == len(rooms)
}
func visitRooms(i int, visited map[int]bool, rooms [][]int) {
	if visited[i] {
		return
	}
	visited[i] = true
	for _, key := range rooms[i] {
		visitRooms(key, visited, rooms)
	}
}
