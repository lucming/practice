package leetcode

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, val := range nums {
		m[val] = true
	}

	longest := 0
	for cur := range m {
		if !m[cur-1] {
			tmp := cur
			curLength := 1
			for m[tmp+1] {
				curLength++
				tmp++
			}
			if curLength > longest {
				longest = curLength
			}
		}
	}

	return longest
}
