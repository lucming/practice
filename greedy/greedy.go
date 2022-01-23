package greedy

import "sort"

//分饼干问题
//g数组描述每个孩子满足值，s数组表示每个饼干的大小
//问题：最多可以让几个孩子满足
//思路1：大饼干分给胃口大的孩子
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	idx := len(s) - 1
	result := 0
	for i := len(g) - 1; i >= 0; i-- {
		if idx >= 0 && s[idx] >= g[i] {
			result++
			idx--
		}
	}
	return result
}

//思路2：小饼干分给胃口小的孩子
func findContentChildren1(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= g[idx] {
			idx++
		}
	}

	return idx
}
