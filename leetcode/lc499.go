package leetcode

func findTargetSumWays(nums []int, target int) int {
	count := 0
	var do func(int, int)
	do = func(start, curSum int) {
		if start == len(nums) {
			if curSum == target {
				count++
			}
			return
		}

		do(start+1, curSum+nums[start])
		do(start+1, curSum-nums[start])
	}

	do(0, 0)
	return count
}
