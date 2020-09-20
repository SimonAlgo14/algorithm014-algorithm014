package main

import "fmt"

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minPathSum(grid [][]int) int {
	if len(grid) < 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row == 0 && col == 0 {
				continue
			} else if row == 0 && col > 0 {
				grid[row][col] = grid[row][col] + grid[row][col-1]
			} else if row > 0 && col == 0 {
				grid[row][col] = grid[row][col] + grid[row-1][col]
			} else if row > 0 && col > 0 {
				grid[row][col] = grid[row][col] + min(grid[row-1][col], grid[row][col-1])
			}
		}
	}

	return grid[rows-1][cols-1]
}

func main() {
	puzzle := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	fmt.Println(minPathSum(puzzle))
}