package greedy

import (
	"sort"
	"strconv"
)

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

//在柠檬水摊上，每一杯柠檬水的售价为 5 美元。
//顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
//每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
//注意，一开始你手头没有任何零钱。
//如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
//示例 1：
//输入：[5,5,5,10,20]
//输出：true
//解释：
//前 3 位顾客那里，我们按顺序收取 3 张 5 美元的钞票。
//第 4 位顾客那里，我们收取一张 10 美元的钞票，并返还 5 美元。
//第 5 位顾客那里，我们找还一张 10 美元的钞票和一张 5 美元的钞票。
//由于所有客户都得到了正确的找零，所以我们输出 true。
//示例 2：
//输入：[5,5,10]
//输出：true
//示例 3：
//输入：[10,10]
//输出：false
//示例 4：
//输入：[5,5,10,10,20]
//输出：false
//解释：
//前 2 位顾客那里，我们按顺序收取 2 张 5 美元的钞票。
//对于接下来的 2 位顾客，我们收取一张 10 美元的钞票，然后返还 5 美元。
//对于最后一位顾客，我们无法退回 15 美元，因为我们现在只有两张 10 美元的钞票。
//由于不是每位顾客都得到了正确的找零，所以答案是 false。
//提示：
//0 <= bills.length <= 10000
//bills[i] 不是 5 就是 10 或是 20
func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			if five < 1 {
				return false
			}
			ten++
			five--
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
				twenty++
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}

//输入：一个数组 points ，其中 points [i] = [xstart,xend] ，
//输出：返回引爆所有气球所必须射出的最小弓箭数。
func findMinArrowShots(points [][]int) int {
	result := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	for i := 1; i < len(points); i++ {
		//没有重叠，需要一剑
		if points[i-1][1] < points[i][0] {
			result++
		} else { //更新最小的右边界
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
//注意: 可以认为区间的终点总是大于它的起点。 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
func eraseOverlapIntervals(intervals [][]int) int {
	result := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	for i := 1; i < len(intervals); i++ {
		//有交集，要执行删除的操作，更新最小右边界
		if intervals[i-1][1] > intervals[i][0] {
			result++
			intervals[i][1] = min(intervals[i-1][1], intervals[i][1])
		}
	}

	return result
}

//给出一个区间的集合，请合并所有重叠的区间。
//示例 1:
//输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//示例 2:
//输入: intervals = [[1,4],[4,5]]
//输出: [[1,5]]
//解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
//提示：
//intervals[i][0] <= intervals[i][1]
func merge(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	result := make([][]int, 0)
	current := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1][1] >= nums[i][0] {
			current[1] = max(nums[i-1][1], nums[i][1])
		} else {
			result = append(result, current)
			current = nums[i]
		}
	}
	result = append(result, current)

	return result
}

//给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
//（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
//示例 1:
//输入: N = 10
//输出: 9
//示例 2:
//输入: N = 1234
//输出: 1234
//示例 3:
//输入: N = 332
//输出: 299
//说明: N 是在 [0, 10^9] 范围内的一个整数。
func monotoneIncreasingDigits(num int) int {
	stringNum := strconv.Itoa(num)
	byteNum := []byte(stringNum)
	if len(byteNum) <= 1 {
		return num
	}

	for i := len(byteNum) - 1; i > 0; i-- {
		if byteNum[i] < byteNum[i-1] {
			byteNum[i-1]--
			for j := i; j < len(byteNum); j++ {
				byteNum[j] = '9'
			}
		}
	}

	result, _ := strconv.Atoi(string(byteNum))
	return result
}

//给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
//你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
//返回获得利润的最大值。
//注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
//示例 1: 输入: prices = [1, 3, 2, 8, 10, 4, 9], fee = 2 输出: 8
//解释: 能够达到的最大利润: 在此处买入 prices[0] = 1 在此处卖出 prices[3] = 8 在此处买入 prices[4] = 4 在此处卖出 prices[5] = 9 总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
func maxProfitHaveFee(prices []int, fee int) int {
	result := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		//买入
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		//未获利，不操作
		if prices[i] >= minPrice && prices[i] <= minPrice+fee {
			continue
		}

		//这边有两层意思：1.获利；2.当前值>前一个节点的值
		if prices[i] > minPrice+fee {
			result += prices[i] - minPrice - fee
			minPrice = prices[i] - fee //已经获利，但是不应该这个时候卖掉，所以最后的获利应该加上递增的这一段的差值
		}
	}

	return result
}
