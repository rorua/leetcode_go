package number_of_islands

func numIslands(grid [][]byte) int {
	n := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				n++
				checkEdges(grid, i, j)
			}
		}
	}
	return n
}

func checkEdges(grid [][]byte, i, j int) {
	if (i < 0 || i >= len(grid)) || (j < 0 || j >= len(grid[i])) || (grid[i][j] == '0') {
		return
	}

	grid[i][j] = '0'
	checkEdges(grid, i, j+1)
	checkEdges(grid, i, j-1)
	checkEdges(grid, i+1, j)
	checkEdges(grid, i-1, j)
}
