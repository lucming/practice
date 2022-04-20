package leetcode

func exist(board [][]byte, word string) bool {
	for i, row := range board {
		for j := range row {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j int, word string, k int) bool {
	if k == len(word) {
		return true
	}

	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[0])-1 {
		return false
	}

	if word[k] != board[i][j] {
		return false
	}

	cur := board[i][j]
	board[i][j] = '0'
	result := dfs(board, i+1, j, word, k+1) ||
		dfs(board, i-1, j, word, k+1) ||
		dfs(board, i, j+1, word, k+1) ||
		dfs(board, i, j-1, word, k+1)
	board[i][j] = cur

	return result
}
