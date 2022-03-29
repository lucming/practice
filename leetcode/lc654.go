package leetcode

import (
	"github.com/lucming/practice/tools"
)

//给一个长度为 N 的数组 nums，其中本来装着 [1..N] 这 N 个元素，无序。
//但是现在出现了一些错误，nums 中的一个元素出现了重复，也就同时导致了另一个元素的缺失。
//请你写一个算法，找到 nums 中的重复元素和缺失元素的值。
//比如说输入：nums = [1,2,2,4]，算法返回 [2,3]。
func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		idx := tools.Abs(nums[i]) - 1
		if nums[idx] < 0 {
			result[0] = tools.Abs(nums[i])
		} else {
			nums[idx] *= -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result[1] = i + 1
		}
	}

	return result
}
