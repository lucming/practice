package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	var do func(bool, int, []int)
	do = func(preUsed bool, idx int, cur []int) {
		if idx == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
			return
		}

		do(false, idx+1, cur)
		if !preUsed && idx > 0 && nums[idx-1] == nums[idx] {
			return
		}
		cur = append(cur, nums[idx])
		do(true, idx+1, cur)
		cur = cur[:len(cur)-1]
	}

	do(false, 0, []int{})
	return result
}
