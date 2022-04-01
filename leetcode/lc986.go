package leetcode

import "github.com/lucming/practice/tools"

func intervalIntersection(a, b [][]int) [][]int {
	result := make([][]int, 0)
	for i, j := 0, 0; i < len(a) && j < len(b); {
		left := tools.Max(a[i][0], b[j][0])
		right := tools.Min(a[i][1], b[j][1])
		if left <= right {
			result = append(result, []int{left, right})
		}
		if a[i][1] < b[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}
