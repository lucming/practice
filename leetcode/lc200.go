package leetcode

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs1(grid, i, j)
			}
		}
	}

	return count
}

func dfs1(grid [][]byte, row, col int) {
	if row >= len(grid) || col >= len(grid[0]) || row < 0 || col < 0 || grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'
	dfs1(grid, row-1, col)
	dfs1(grid, row+1, col)
	dfs1(grid, row, col-1)
	dfs1(grid, row, col+1)
}
