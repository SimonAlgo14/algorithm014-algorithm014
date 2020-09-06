package main


func updateBoard(board [][]byte, click []int) [][]byte {
	r, c := click[0], click[1]
	if board[r][c] == 'M' {
		// 1. If a mine ('M') is revealed, then the game is over - change it to
		// 'X'.
		board[r][c] = 'X'
		return board
	} else if board[r][c] == 'E' {
		adjSquares := adjacentSquares(board, click)
		var numAdjMines byte
		for _, adj := range adjSquares {
			r, c := adj[0], adj[1]
			if board[r][c] == 'M' {
				numAdjMines++
			}
		}
		if numAdjMines == 0 {
			// 2. If an empty square ('E') with no adjacent mines is revealed, then
			// change it to revealed blank ('B') and all of its adjacent
			// unrevealed squares should be revealed recursively.
			board[r][c] = 'B'
			for _, adj := range adjSquares {
				r, c := adj[0], adj[1]
				if board[r][c] == 'M' || board[r][c] == 'E' {
					updateBoard(board, adj)
				}
			}
		} else {
			// 3. If an empty square ('E') with at least one adjacent mine is
			// revealed, then change it to a digit ('1' to '8') representing the
			// number of adjacent mines.
			board[r][c] = '0' + numAdjMines
		}
	}
	// 4. Return the board when no more squares will be revealed.
	return board
}

func adjacentSquares(board [][]byte, click []int) [][]int {
	nr := len(board)
	nc := len(board[0])
	var adjs [][]int

	appendAdj := func(adj []int) {
		if adj[0] >= 0 && adj[0] < nr && adj[1] >= 0 && adj[1] < nc {
			adjs = append(adjs, adj)
		}
	}

	r, c := click[0], click[1]
	appendAdj([]int{r - 1, c - 1})
	appendAdj([]int{r - 1, c})
	appendAdj([]int{r - 1, c + 1})
	appendAdj([]int{r, c - 1})
	appendAdj([]int{r, c + 1})
	appendAdj([]int{r + 1, c - 1})
	appendAdj([]int{r + 1, c})
	appendAdj([]int{r + 1, c + 1})

	return adjs
}
