package leetcode

func subarraySum(nums []int, k int) int {
	count := 0
	pre := 0
	m := make(map[int]int)
	m[0] = 1

	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}

	return count
}
