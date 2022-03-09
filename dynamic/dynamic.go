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

//打家劫舍问题
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//示例 1： 输入：[1,2,3,1] 输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。   偷窃到的最高金额 = 1 + 3 = 4 。
//示例 2： 输入：[2,7,9,3,1] 输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。   偷窃到的最高金额 = 2 + 9 + 1 = 12 。
func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return tools.Max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = tools.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		// 第i个房间不偷 dp[i-1]
		// 第i个房间偷 dp[i-2]+nums[i] 那么i-1就一定不偷
		dp[i] = tools.Max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

//打家劫舍问题
//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。
//示例 1：
//输入：nums = [2,3,2] 输出：3
//解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
//示例 2： 输入：nums = [1,2,3,1] 输出：4
//解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。偷窃到的最高金额 = 1 + 3 = 4 。
//示例 3： 输入：nums = [0] 输出：0
func robLimitStartAndEnd(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = tools.Max(dp[start], dp[start+1])
	for i := start + 2; i <= end; i++ {
		dp[i] = tools.Max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[end]
}

func rob1(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	val1 := robLimitStartAndEnd(nums, 0, len(nums)-2)
	val2 := robLimitStartAndEnd(nums, 1, len(nums)-1)

	return tools.Max(val1, val2)
}

//给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//示例 1：
//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//示例 2：
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

//1. 暴力，两层循环，找到结束和开始差值最大，即为最大利润
func maxProfit(prices []int) int {
	result := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			result = tools.Max(result, prices[j]-prices[i])
		}
	}

	return result
}

//2. 贪心，一层循环，分别记录最小和最大的价格，做差即为最大利润
func maxProfit1(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	result := 0
	low := prices[0]
	for i := 0; i < len(prices); i++ {
		//买入最小的价格
		low = tools.Min(low, prices[i])
		//取最大的差值
		result = tools.Max(result, prices[i]-low)
	}

	return result
}

//动态规划的写法1
//dp[i][0] 代表第i天持有 （第i天买入，或者i天之前已经买入，持续第i-1天的状态即可）
//dp[i][1] 代表第i天不持有 （第i天卖出，或者昨天已经卖出）
func maxProfit2(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = tools.Max(dp[i-1][0], -prices[i])
		dp[i][1] = tools.Max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}

func maxProfit3(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := make([]int, 2)
	// 记录一次交易，一次交易有买入卖出两种状态
	// 0代表持有，1代表卖出
	dp[0] = -prices[0]
	dp[1] = 0
	for i := 1; i < len(prices); i++ {
		// 前一天持有；或当天买入
		dp[0] = tools.Max(dp[0], -prices[i-1])
		// 前一天卖出；或当天卖出, 当天要卖出，得前一天持有才行
		dp[1] = tools.Max(dp[1], dp[0]+prices[i-1])
	}

	return dp[1]
}

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
//示例 1： 输入：nums = [10,9,2,5,3,7,101,18] 输出：4 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//示例 2： 输入：nums = [0,1,0,3,2,3] 输出：4
//示例 3： 输入：nums = [7,7,7,7,7,7,7] 输出：1

//思路：dp数组每个val可以先初始化为1,开始遍历，
//内层循环逻辑：如果当前val>前面遍历过的某个值，说明当前值可以和前面的值组成递增子数组，dp[i]=dp[j]+1
//其实就是找前面比当前值小的最长序列，当前最长就是前面最长+1
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	result := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = tools.Max(dp[i], dp[j]+1)
			}
		}
		if result < dp[i] {
			result = dp[i]
		}
	}

	return result
}

//最长连续子数组
//给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
//示例 1： 输入：nums = [1,3,5,4,7] 输出：3 解释：最长连续递增序列是 [1,3,5], 长度为3。 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为 5 和 7 在原数组里被 4 隔开。
//示例 2： 输入：nums = [2,2,2,2,2] 输出：1 解释：最长连续递增序列是 [2], 长度为1。
// if nums[i]>nums[i-1] dp[i]=dp[i-1]+1
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	dp[0] = 1
	maxLength := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}

	return maxLength
}

//最长重复子数组长度
//给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
//示例：
//输入： A: [1,2,3,2,1] B: [3,2,1,4,7] 输出：3 解释： 长度最长的公共子数组是 [3, 2, 1] 。
func findLength(A, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}

	dp := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		dp[i] = make([]int, len(B)+1)
	}
	maxLength := 0
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if maxLength < dp[i][j] {
				maxLength = dp[i][j]
			}
		}
	}

	return maxLength
}

func findLength1(A, B []int) int {
	dp := make([]int, len(B)+1)
	maxLength := 0
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[j] = dp[j-1] + 1
			}
			if maxLength < dp[j] {
				maxLength = dp[j]
			}
		}
	}

	return maxLength
}

//最长公共子序列
//给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
//一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
//若这两个字符串没有公共子序列，则返回 0。
//示例 1:
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace"，它的长度为 3。
//示例 2:
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc"，它的长度为 3。
//示例 3:
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0。
func longestCommonSubsequence(text1, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = tools.Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

//最大连续子数组和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//示例: 输入: [-2,1,-3,4,-1,2,1,-5,4] 输出: 6 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//dp[i]=max(dp[i-1]+nums[i],nums[i])
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := 0
	for i := 1; i < len(nums); i++ {
		dp[i] = tools.Max(dp[i-1]+nums[i], nums[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}

	return maxSum
}

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//示例 1： 输入：s = "abc", t = "ahbgdc" 输出：true
//示例 2： 输入：s = "axc", t = "ahbgdc" 输出：false
func isSubsequence(s, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}

//不同的子序列
//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
func numDistinct(s, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(s)][len(t)]
}
