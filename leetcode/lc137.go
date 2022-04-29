package leetcode

//给你一个整数数组nums，除某个元素仅出现一次外，其余每个元素都恰出现三次 。请你找出并返回那个只出现了一次的元素。
//map记录每个数字出现的次数 时间复杂度：O(n) 空间复杂度：O(n)
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}

func singleNumber1(nums []int) int {
	sum := 0
	s := make(map[int]struct{})
	for _, num := range nums {
		sum += num
		s[num] = struct{}{}
	}

	sumWrapper := 0
	for num := range s {
		sumWrapper += num
	}

	return (3*sumWrapper - sum) / 2
}
