package permutationAndCombination

//题目：给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合
// 思路：每次从集合中选取元素，可选择的范围随着选择的进行而收缩，调整可选择的范围。可以发现n相当于树的宽度，k相当于树的深度。
func getCombinations(n int, k int) [][]int {
	result := make([][]int, 0)
	if n <= 0 || k <= 0 || n < k {
		return result
	}

	var do func(int, int, int, []int)
	do = func(n int, k int, start int, track []int) {
		if len(track) == k {
			tmp := make([]int, k)
			copy(tmp, track)
			result = append(result, tmp)
		}
		if len(track)+n-start+1 < k {
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			do(n, k, i+1, track)
			track = track[:len(track)-1]
		}
	}

	do(n, k, 1, []int{})
	return result
}
