package main

func numIslands(grid [][]byte) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		cols := len(grid[0])

		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count ++
				mask(grid, row, col)
			}
		}
	}

	return count
}

func mask(grid [][]byte, row int, col int) {
	rows := len(grid)
	cols := len(grid[0])

	if row >= rows || col >= cols {
		return
	}

	if col + 1 < cols && grid[row][col+1] == '1' {
		grid[row][col+1] = '0'
		mask(grid, row, col + 1)
	}

	if col - 1 >= 0 && grid[row][col-1] == '1' {
		grid[row][col-1] = '0'
		mask(grid, row, col - 1)
	}

	if row + 1 < rows && grid[row+1][col] == '1' {
		grid[row+1][col] = '0'
		mask(grid, row+1, col)
	}

	if row - 1 >= 0 && grid[row-1][col] == '1' {
		grid[row-1][col] = '0'
		mask(grid, row-1, col)
	}
}


/*
  1  1  1
  0  1  0
  1  1  1
 */
