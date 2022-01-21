package permutationAndCombination

import (
	"sort"
	"strconv"
	"strings"
)

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

//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//返回 s 所有可能的分割方案。
//示例: 输入: "aab" 输出: [ ["aa","b"], ["a","a","b"] ]
func partition(s string) [][]string {
	var tmpStr []string
	var result [][]string
	backTracing(s, 0, tmpStr, &result)
	return result
}

func backTracing(s string, start int, tmpStr []string, result *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(tmpStr))
		copy(tmp, tmpStr)
		*result = append(*result, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if !isCycle(s, start, i) {
			continue
		}
		tmpStr = append(tmpStr, s[start:i+1])
		backTracing(s, i+1, tmpStr, result)
		tmpStr = tmpStr[:len(tmpStr)-1]
	}
}

func isCycle(s string, start, end int) bool {
	l, r := start, end
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
func RestoreIp(s string) []string {
	var path, result []string
	back(s, 0, path, &result)
	return result
}

func back(s string, start int, path []string, result *[]string) {
	if start == len(s) && len(path) == 4 {
		ip := strings.Join(path, ".")
		*result = append(*result, ip)
		return
	}

	for i := start; i < len(s); i++ {
		path = append(path, s[start:i+1])
		if len(path) <= 4 && i-start+1 <= 3 && isIp(s, start, i) {
			back(s, i+1, path, result)
		} else {
			return
		}
		path = path[:len(path)-1]
	}
}

func isIp(s string, start, end int) bool {
	i, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if i > 255 {
		return false
	}
	return true
}

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//说明：解集不能包含重复的子集。
//示例: 输入: nums = [1,2,3] 输出: [[3],[1],[2],[1,2,3],[1,3],[2,3],[1,2],[]]
func Subset(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	var do func([]int, []int, int)
	do = func(nums []int, cur []int, start int) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			do(nums, cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}

	do(nums, []int{}, 0)
	return result
}

//给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
//示例:
//输入: [4, 6, 7, 7]
//输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
//说明:
//给定数组的长度不会超过15。
//数组中的整数范围是 [-100,100]。
//给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	var do func([]int, int, []int)
	do = func(nums []int, start int, cur []int) {
		if len(cur) > 1 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
		}
		//用于标记历史数据是否出现过
		history := make([]bool, 201)
		for i := start; i < len(nums); i++ {
			//两种情况：
			//  1. 当前元素已经出现过，继续找
			//  2. 当前元素<子集最后一个元素，不必继续
			if history[nums[i]+100] == true || len(cur) > 0 && nums[i] < cur[len(cur)-1] {
				continue
			}
			history[nums[i]+100] = true
			cur = append(cur, nums[i])
			do(nums, i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}

	do(nums, 0, []int{})
	return result
}

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	history := make([]bool, len(nums))
	var do func([]int, int, []int)
	do = func(nums []int, lenNums int, cur []int) {
		if len(cur) == lenNums {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
		}
		for i := 0; i < lenNums; i++ {
			if history[i] {
				continue
			}
			cur = append(cur, nums[i])
			history[i] = true
			do(nums, lenNums, cur)
			cur = cur[:len(cur)-1]
			history[i] = false
		}
	}
	do(nums, len(nums), []int{})
	return result
}

func permute1(nums []int) [][]int {
	result := make([][]int, 0)
	var do func([]int, int, []int)
	do = func(nums []int, lenNums int, cur []int) {
		if len(nums) == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
		}
		for i := 0; i < lenNums; i++ {
			used := nums[i]
			cur = append(cur, used)
			nums = append(nums[:i], nums[i+1:]...)
			do(nums, len(nums), cur)
			cur = cur[:len(cur)-1]
			nums = append(nums[:i], append([]int{used}, nums[i:]...)...)
		}
	}
	do(nums, len(nums), []int{})
	return result
}

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//1 <= nums.length <= 8
//-10 <= nums[i] <= 10
//示例 1：
//输入：nums = [1,1,2]
//输出： [[1,1,2], [1,2,1], [2,1,1]]
//示例 2：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permuteHasSameNum(nums []int) [][]int {
	result := make([][]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var do func([]int, []int)
	do = func(nums []int, cur []int) {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
		}
		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}

			cur = append(cur, nums[i])
			used[i] = true
			do(nums, cur)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	do(nums, []int{})
	return result
}

func permuteHasSameNum1(nums []int) [][]int {
	result := make([][]int, 0)
	var do func([]int, int, []int)
	do = func(nums []int, lenNums int, cur []int) {
		if len(nums) == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			result = append(result, tmp)
		}
		used := make([]bool, 21)
		for i := 0; i < lenNums; i++ {
			if used[nums[i]+10] {
				continue
			}

			d := nums[i]
			cur = append(cur, d)
			used[nums[i]+10] = true
			nums = append(nums[:i], nums[i+1:]...)
			do(nums, len(nums), cur)
			cur = cur[:len(cur)-1]
			nums = append(nums[:i], append([]int{d}, nums[i:]...)...)
		}
	}

	do(nums, len(nums), []int{})
	return result
}
