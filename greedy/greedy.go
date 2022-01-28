package greedy

import "sort"

//分饼干问题
//g数组描述每个孩子满足值，s数组表示每个饼干的大小
//问题：最多可以让几个孩子满足
//思路1：大饼干分给胃口大的孩子
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	idx := len(s) - 1
	result := 0
	for i := len(g) - 1; i >= 0; i-- {
		if idx >= 0 && s[idx] >= g[i] {
			result++
			idx--
		}
	}
	return result
}

//思路2：小饼干分给胃口小的孩子
func findContentChildren1(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= g[idx] {
			idx++
		}
	}

	return idx
}

//摆动序列
//如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。
//例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和 [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
//给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
//示例 1:
//输入: [1,7,4,9,2,5]
//输出: 6
//解释: 整个序列均为摆动序列。
//示例 2:
//输入: [1,17,5,10,13,15,10,5,16,8]
//输出: 7
//解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。
//示例 3:
//输入: [1,2,3,4,5,6,7,8,9]
//输出: 2
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	curDiff, preDiff := 0, 0
	result := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if curDiff > 0 && preDiff <= 0 || curDiff <= 0 && preDiff > 0 {
			result++
			preDiff = curDiff
		}
	}

	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//动态规划
func wiggleMaxLength1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([][]int, 1005)
	for i := 0; i < 1005; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 1, 1

	for i := 1; i < len(nums); i++ {
		dp[i][0], dp[i][1] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			}
		}
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

//最大连续子数组和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//示例: 输入: [-2,1,-3,4,-1,2,1,-5,4] 输出: 6 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
//暴力
func maxSubArray(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			result = max(result, count)
		}
	}
	return result
}

//贪心算法（eg:[-2,1]这种情况肯定会从1开始算起，因为负数会拉低总和），
//局部最优（如果当前连续和为负数的时候立刻放弃，从下一个下标开始重新计算连续和，因为负数只会拉低最后的结果）
func maxSubArray1(nums []int) int {
	result := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count < 0 {
			count = 0
			continue
		}
		result = max(result, count)
	}

	return result
}

//动态规划（每次会取前面最大和+当前值和当前值的较大值）
func maxSubArray2(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}

	return result
}

//买卖股票的问题
//贪心算法，每次只选择当前有收益的进行累计
func maxProfit(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	result := 0
	for i := 1; i < len(nums); i++ {
		profit := nums[i] - nums[i-1]
		if profit > 0 {
			result += profit
		}
	}

	return result
}

//动态规划的解法
func maxProfit1(nums []int) int {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] -= nums[0]
	for i := 1; i < len(nums); i++ {
		//第i天持股所剩最多现金=max(第i-1天持股票所剩的现金,第i-1天持现金-今天买入的现金)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-nums[i])
		//第i天持有的最多现金=max(第i-1天持有最多的现金,第i-1天的股票最多现金+第i天卖出去的股票)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+nums[i])
	}

	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

//给定一个非负整数数组，你最初位于数组的第一个位置。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//判断你是否能够到达最后一个位置。
//示例 1:
//输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
//示例 2:
//输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(nums[i]+i, cover)
		if cover >= len(nums) {
			return true
		}
	}

	return false
}

//题目：给定数组，判断最少几步可以走到结束，数组中每个元素代表可以最多跳的步数
//示例:
//输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//说明: 假设你总是可以到达数组的最后一个位置。
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	curRange, nextRange := 0, 0
	steps := 0

	for i := 0; i < len(nums); i++ {
		nextRange = max(nextRange, nums[i]+i)
		if i != curRange {
			continue
		}
		if curRange == len(nums)-1 {
			continue
		}
		steps++
		curRange = nextRange
		if nextRange >= len(nums)-1 {
			break
		}
	}

	return steps
}

func jump1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	cur, next := 0, 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		next = max(next, nums[i]+i)
		if i != cur {
			continue
		}
		step++
		cur = next
	}

	return step
}

//分发糖果
//老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
//你需要按照以下要求，帮助老师给这些孩子分发糖果：
//每个孩子至少分配到 1 个糖果。
//相邻的孩子中，评分高的孩子必须获得更多的糖果。
//那么这样下来，老师至少需要准备多少颗糖果呢？
//示例 1:
//输入: [1,0,2]
//输出: 5
//解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
//示例 2:
//输入: [1,2,2]
//输出: 4
//解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
func candy(arr []int) int {
	historys := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		historys[i] = 1
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			historys[i] = historys[i-1] + 1
		}
	}
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			historys[i-1] = max(historys[i]+1, historys[i-1])
		}
	}
	sum := 0
	for i := 0; i < len(historys); i++ {
		sum += historys[i]
	}
	return sum
}
