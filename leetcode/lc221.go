package leetcode

import "github.com/lucming/practice/tools"

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if dp[i][j] == 0 {
				continue
			}
			dp[i][j] = tools.Min(dp[i-1][j], tools.Min(dp[i-1][j-1], dp[i][j-1])) + 1
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}

	return maxSide * maxSide
}
