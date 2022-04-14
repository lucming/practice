package leetcode

import "github.com/lucming/practice/tools"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	result := 0
	right := 0
	for left := 0; left < len(s); left++ {
		if left != 0 {
			delete(m, s[left-1])
		}
		for j := right; j < len(s); j++ {
			if m[s[j]] != 0 {
				break
			}
			m[s[j]]++
			right++
		}

		result = tools.Max(result, right-left)
	}

	return result
}
