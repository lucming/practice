package leetcode

import "github.com/lucming/practice/tools"

//问题：找到数组中缺少的那个数字
//思路1:下标和value对应，然后一路异或
//思路2：等差数列求和，原本的总和-当前的总和=差的数字
func missingNumber(nums []int) int {
	newIdx := len(nums)
	result := 0
	result ^= newIdx
	for i := 0; i < len(nums); i++ {
		result ^= i ^ nums[i]
	}

	return result
}

func missingNumber1(nums []int) int {
	sum := tools.Sum(nums)
	specicalSum := len(nums) * (len(nums) + 1) / 2

	return specicalSum - sum
}
