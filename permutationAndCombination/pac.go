package permutationAndCombination

import "sort"

//题目：给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合
// 思路：每次从集合中选取元素，可选择的范围随着选择的进行而收缩，调整可选择的范围。可以发现n相当于树的宽度，k相当于树的深度。
func getCombinations(n int, k int) [][]int {
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

//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//candidates 中的数字可以无限制重复被选取。
//说明：
//所有数字（包括 target）都是正整数。
//解集不能包含重复的组合。
//示例 1： 输入：candidates = [2,3,6,7], target = 7, 所求解集为： [[7],[2,2,3]]
//示例 2： 输入：candidates = [2,3,5], target = 8, 所求解集为： [[2,2,2,2],[2,3,3],[3,5]]
func CombinationSum(condidates []int, target int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	backTrack(0, 0, target, condidates, track, &result)
	return result
}

func backTrack(start, sum, target int, condidates []int, track []int, result *[][]int) {
	if sum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(condidates); i++ {
		track = append(track, condidates[i])
		sum += condidates[i]
		backTrack(i, sum, target, condidates, track, result)
		track = track[:len(track)-1]
		sum -= condidates[i]
	}
}

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//candidates 中的每个数字在每个组合中只能使用一次。
//说明： 所有数字（包括目标数）都是正整数。 解集不能包含重复的组合。
//示例 1: 输入: candidates = [10,1,2,7,6,1,5], target = 8, 所求解集为: [[1,7],[1,2,5],[2,6],[1,1,6]]
//示例 2: 输入: candidates = [2,5,2,1,2], target = 5, 所求解集为: [[1,2,2],[5]]
func CombinationSum1(condidates []int, target int) [][]int {
	result := make([][]int, 0)
	track := make([]int, 0)
	used := make(map[int]bool)
	sort.Ints(condidates)
	backTrack1(0, 0, target, condidates, track, &result, used)
	return result
}

func backTrack1(start, sum, target int, condidates []int, track []int, result *[][]int, used map[int]bool) {
	if sum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*result = append(*result, tmp)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(condidates); i++ {
		if i > 0 && condidates[i] == condidates[i-1] && used[i-1] == false {
			continue
		}
		track = append(track, condidates[i])
		sum += condidates[i]
		used[i] = true
		backTrack1(i+1, sum, target, condidates, track, result, used)
		track = track[:len(track)-1]
		sum -= condidates[i]
		used[i] = false
	}
}
