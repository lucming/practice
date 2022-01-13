package permutationAndCombination

//题目：给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合
func getCombinations(n, k int) [][]int {
	result := make([][]int, 0)
	if n <= 0 || k <= 0 || k > n {
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

//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
//说明：
//所有数字都是正整数。
//解集不能包含重复的组合。
//示例 1: 输入: k = 3, n = 7 输出: [[1,2,4]]
//示例 2: 输入: k = 3, n = 9 输出: [[1,2,6], [1,3,5], [2,3,4]]
func CombinationsSum(k, n int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	if n <= 0 || k <= 0 || n < k {
		return result
	}
	var do func(int, int, int, []int)
	do = func(n, k, start int, track []int) {
		if len(track) == k {
			sum := 0
			tmp := make([]int, k)
			for k, v := range track {
				sum += v
				tmp[k] = v
			}
			if sum == n {
				result = append(result, tmp)
			}
			return
		}

		for i := start; i <= 9-(k-len(track))+1; i++ {
			track = append(track, i)
			do(n, k, i+1, track)
			track = track[:len(track)-1]
		}
	}

	do(n, k, 1, track)
	return result
}
