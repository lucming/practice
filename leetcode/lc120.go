package leetcode

import "github.com/lucming/practice/tools"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		//在最左遍，当前值只能由前一行的最左边得出
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < len(triangle[i]); j++ {
			dp[i][j] = tools.Min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		//在最右边，当前值只能由前一行的最后一个值推出
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	// 最小的路径在最后一行
	result := dp[n-1][0]
	for i := 0; i < n; i++ {
		if result > dp[n-1][i] {
			result = dp[n-1][i]
		}
	}

	return result
}
