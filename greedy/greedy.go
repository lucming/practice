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
