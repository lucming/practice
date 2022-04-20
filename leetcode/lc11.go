package leetcode

import "github.com/lucming/practice/tools"

func maxArea(height []int) int {
	maxValue := 0
	l, r := 0, len(height)-1
	for l <= r {
		cur := (r - l) * tools.Min(height[l], height[r])
		if maxValue < cur {
			maxValue = cur
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxValue
}
