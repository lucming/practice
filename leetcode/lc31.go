package leetcode

import "github.com/lucming/practice/tools"

//下一个全排列
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	tools.Reverse(nums[i+1:])
}
