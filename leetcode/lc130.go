package leetcode

func solve(board [][]byte) {
	//上下边界
	for i := 0; i < len(board); i++ {
		dfs2(board, i, 0)
		dfs2(board, i, len(board[0])-1)
	}
	//左右边界
	for i := 0; i < len(board[0]); i++ {
		dfs2(board, 0, i)
		dfs2(board, len(board)-1, i)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs2(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}

	board[x][y] = 'A'
	dfs2(board, x, y+1)
	dfs2(board, x, y-1)
	dfs2(board, x-1, y)
	dfs2(board, x+1, y)
}
