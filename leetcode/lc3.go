package leetcode

import "github.com/lucming/practice/tools"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	result := 0
	right := 0
	for idx := 0; idx < len(s); idx++ {
		if idx != 0 {
			delete(m, s[idx-1])
		}
		for j := right; j < len(s); j++ {
			if m[s[j]] != 0 {
				break
			}
			m[s[j]]++
			right++
		}

		result = tools.Max(result, right-idx)
	}

	return result
}
