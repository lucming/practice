package leetcode

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	var do func(int, []int)
	do = func(start int, cur []int) {
		if start > n {
			return
		}
		if len(cur) == k {
			tmp := make([]int, k)
			copy(tmp, cur)
			result = append(result, tmp)
		}

		cur = append(cur, start)
		do(start+1, cur)
		cur = cur[:len(cur)-1]
	}

	do(1, []int{})
	return result
}
