package leetcode

import "sort"

func frequencySort(s string) string {
	maxCount := 0
	valExistsMap := make(map[byte]int, 0)
	for idx := range s {
		valExistsMap[s[idx]]++
		if maxCount < valExistsMap[s[idx]] {
			maxCount = valExistsMap[s[idx]]
		}
	}

	buckets := make([][]byte, maxCount+1)
	for k, v := range valExistsMap {
		buckets[v] = append(buckets[v], k)
	}

	result := make([]byte, 0, len(s))
	for i := maxCount; i > 0; i-- {
		sort.Slice(buckets[i], func(m, n int) bool {
			return buckets[i][m] < buckets[i][n]
		})
		for _, cur := range buckets[i] {
			tmp := make([]byte, 0)
			for j := 0; j < i; j++ {
				tmp = append(tmp, cur)
			}
			result = append(result, tmp...)
		}
	}

	return string(result)
}
