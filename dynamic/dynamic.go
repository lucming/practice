package dynamic

import "github.com/lucming/practice/tools"

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
