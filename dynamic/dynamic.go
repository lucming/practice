package dynamic

import (
	"github.com/lucming/practice/tools"
)

//斐波那契数列
//f(n)=f(n-1)+f(n-2)
func fib(n int) int {
	if n < 2 {
		return n
	}

	morepre, pre, current := 0, 1, 0
	for i := 2; i <= n; i++ {
		current = pre + morepre
		morepre = pre
		pre = current
	}

	return current
}

//花费最小体力爬楼梯
func minCostClimbingStairs(cost []int) int {
	morepre, pre, current := cost[0], cost[1], 0
	for i := 2; i < len(cost); i++ {
		current = tools.Min(morepre, pre) + cost[i]
		morepre = pre
		pre = current
	}

	return tools.Min(morepre, pre)
}

func minCostClimbingStairs1(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = tools.Min(dp[i-2], dp[i-1]) + cost[i]
	}

	return tools.Min(dp[len(dp)-1], dp[len(dp)-2])
}

//一个机器人位于一个 m x n 网格的左上角
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角
//问总共有多少条不同的路径？
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

//一个机器人位于一个 m x n 网格的左上角
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
func uniquePathsHaveStone(girds [][]int) int {
	dp := make([][]int, len(girds))
	for i := 0; i < len(girds); i++ {
		dp[i] = make([]int, len(girds[i]))
		for j := 0; j < len(girds[i]); j++ {
			if (i == 0 || j == 0) && girds[i][j] != 1 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < len(girds); i++ {
		for j := 1; j < len(girds[i]); j++ {
			if girds[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			dp[i] = tools.Max(dp[i], tools.Max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}
