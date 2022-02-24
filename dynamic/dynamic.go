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

//有n件物品和一个最多能背重量为w 的背包。
//第i件物品的重量是weight[i]，得到的价值是value[i] 。
//每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。
func bagProblem01(weight, value []int, bagweight int) int {
	dp := make([][]int, len(weight))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bagweight+1)
	}

	//初始化
	//d[0][j] 在下标为0的物品中选择，放进容量为j的背包，价值的最大值
	for j := weight[0]; j <= bagweight; j++ {
		dp[0][j] += value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := weight[i]; j <= bagweight; j++ {
			//dp[i][j] 表示从下标为[0-i]的物品里任意取，放进容量为j的背包，价值总和最大是多少。
			dp[i][j] = tools.Max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}

	return dp[len(weight)-1][bagweight]
}

//一维数组
func bagProblem(weight, value []int, bagweight int) int {
	dp := make([]int, bagweight+1)
	for i := 0; i < len(weight); i++ {
		for j := bagweight; j >= weight[i]; j-- {
			dp[j] = tools.Max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	return dp[bagweight]
}

//给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//示例 1: 输入: [1, 5, 11, 5] 输出: true 解释: 数组可以分割成 [1, 5, 5] 和 [11].
//示例 2: 输入: [1, 2, 3, 5] 输出: false 解释: 数组不能分割成两个元素和相等的子集.
func canPartition(nums []int) bool {
	sum := tools.Sum(nums)
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = tools.Max(dp[j], dp[j-num]+num)
		}
	}

	return dp[target] == target
}

//有一堆石头，每块石头的重量都是正整数。
//每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//如果 x == y，那么两块石头都会被完全粉碎； 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。 最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回 0。
//示例：
//输入：[2,7,4,1,8,1]
//输出：1
//解释： 组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]， 组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]， 组合 2 和 1，得到 1，所以数组转化为 [1,1,1]， 组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
func lastStoneWeight(stones []int) int {
	sumStone := tools.Sum(stones)
	target := sumStone / 2
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = tools.Max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sumStone - 2*dp[target]
}

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
//返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
//示例：
//输入：nums: [1, 1, 1, 1, 1], S: 3
//输出：5
//解释：
//-1+1+1+1+1 = 3
//+1-1+1+1+1 = 3
//+1+1-1+1+1 = 3
//+1+1+1-1+1 = 3
//+1+1+1+1-1 = 3
//一共有5种方法让最终目标和为3。
func findTargetSumWays(nums []int, target int) int {
	sum := tools.Sum(nums)
	if target > sum {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[bag]
}

//给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
//示例 1: 输入: amount = 5, coins = [1, 2, 5] 输出: 4 解释: 有四种方式可以凑成总金额: 5=5 5=2+2+1 5=2+1+1+1 5=1+1+1+1+1
//示例 2: 输入: amount = 3, coins = [2] 输出: 0 解释: 只用面额2的硬币不能凑成总金额3。
//示例 3: 输入: amount = 10, coins = [10] 输出: 1
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	return dp[amount]
}

//给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
//示例:
//nums = [1, 2, 3] target = 4
//所有可能的组合为： (1, 1, 1, 1) (1, 1, 2) (1, 2, 1) (1, 3) (2, 1, 1) (2, 2) (3, 1)
//因此输出为 7。
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}

	return dp[target]
}

//爬楼梯
func claimstairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for j := 1; j <= n; j++ { //遍历背包
		for i := 1; i <= 2; i++ { //遍历物品
			if j >= i {
				dp[j] += dp[j-i]
			}
		}
	}

	return dp[n]
}
