package binarySearch

import "github.com/lucming/practice/tools"

//偷香蕉的问题
//珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。
//珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
//珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
//返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
//示例 1：
//输入: piles = [3,6,7,11], H = 8
//输出: 4
//示例 2：
//输入: piles = [30,11,23,4,20], H = 5
//输出: 30
//示例 3：
//输入: piles = [30,11,23,4,20], H = 6
//输出: 23

//分析：分析吃香蕉速度的最大值和最小值，min:1 max:max(piles)
//所以可以顺序判断每个速度吃掉香蕉需要的时间，如果还好是h，直接返回
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, tools.MaxValue(piles)+1
	for left < right {
		mid := (left + right) / 2
		if canFinish(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canFinish(piles []int, h, speed int) bool {
	nums := 0
	for i := 0; i < len(piles); i++ {
		cur := piles[i] / speed
		if piles[i]%speed > 0 {
			cur += 1
		}

		nums += cur
	}

	return nums <= h
}

//传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
//传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
//返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。
//输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
//输出：15
//解释：
//船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
//第 1 天：1, 2, 3, 4, 5
//第 2 天：6, 7
//第 3 天：8
//第 4 天：9
//第 5 天：10
//请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。
//
//示例 2：
//输入：weights = [3,2,2,4,1,4], days = 3
//输出：6
//解释：
//船舶最低载重 6 就能够在 3 天内送达所有包裹，如下所示：
//第 1 天：3, 2
//第 2 天：2, 4
//第 3 天：1, 4
//
//示例 3：
//输入：weights = [1,2,3,1,1], days = 4
//输出：3
//解释：
//第 1 天：1
//第 2 天：2
//第 3 天：3
//第 4 天：1, 1

//分析：船的承载能力 min:max(weights), max:sum(weights)
func shipWithinDays(weights []int, days int) int {
	left, right := tools.MaxValue(weights), tools.Sum(weights)+1
	for left < right {
		mid := (left + right) / 2
		if canFinish1(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

//判断载重为cap能否在days天运完
func canFinish1(weights []int, days, cap int) bool {
	i := 0
	for day := 0; day < days; day++ {
		for maxCap := cap - weights[i]; maxCap >= 0; maxCap -= weights[i] {
			i++
			if i == len(weights) {
				return true
			}
		}
	}

	return false
}

//判断子序列
func isSubsequence(s, t string) bool {
	lenS, lenT := len(s), len(t)
	i, j := 0, 0
	for i < lenS && j < lenT {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == lenS
}
